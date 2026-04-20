package config

import (
	"fmt"
	"io"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/goccy/go-yaml"
)

// Parse reads a YAML document and returns a populated Config.
//
// env provides the source for $VAR / ${VAR} substitution. Callers that want
// process env should pass envFromOSEnviron() (or use ParseFile, which does).
func Parse(r io.Reader, env map[string]string) (*Config, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return parseBytes(data, env)
}

// ParseFile is a convenience wrapper that reads `path`, uses os.Environ as
// the substitution source, and returns the directory of the YAML file so
// callers can resolve entries_from_files globs against it.
func ParseFile(path string) (*Config, string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, "", err
	}
	cfg, err := parseBytes(data, envFromOSEnviron())
	if err != nil {
		return nil, "", err
	}
	return cfg, filepath.Dir(path), nil
}

func parseBytes(data []byte, env map[string]string) (*Config, error) {
	substituted, err := substituteEnv(data, env)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.UnmarshalWithOptions(substituted, &cfg, yaml.Strict()); err != nil {
		return nil, fmt.Errorf("config: parse YAML: %w", err)
	}

	if cfg.Version != 0 && cfg.Version != 1 {
		return nil, fmt.Errorf("config: unsupported version %d (only 1 is supported)", cfg.Version)
	}

	if err := populateExternalIDs(&cfg); err != nil {
		return nil, err
	}
	if err := validateRefs(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// substituteEnv replaces $VAR and ${VAR} occurrences in data using env. It
// returns an error listing every unset variable (not just the first) so the
// user can fix all of them in one round-trip.
//
// Uses os.Expand for the substitution pass — same syntax rules as shell
// parameter expansion ($name and ${name}, name = [A-Za-z_][A-Za-z0-9_]*).
func substituteEnv(data []byte, env map[string]string) ([]byte, error) {
	missing := map[string]struct{}{}
	out := os.Expand(string(data), func(name string) string {
		if v, ok := env[name]; ok {
			return v
		}
		missing[name] = struct{}{}
		return "${" + name + "}" // preserve the literal so the error message can point at it
	})
	if len(missing) > 0 {
		names := slices.Sorted(maps.Keys(missing))
		return nil, fmt.Errorf("config: unset environment variable(s): %s", strings.Join(names, ", "))
	}
	return []byte(out), nil
}

// populateExternalIDs walks the parsed Config and fills ExternalID / Key
// fields from their containing map keys. Also rejects empty external_ids
// (the map key is the external_id; an empty key is nonsense).
func populateExternalIDs(cfg *Config) error {
	for k, ts := range cfg.ToolSets {
		if k == "" {
			return fmt.Errorf("config: tool_sets: empty external_id key")
		}
		if ts == nil {
			return fmt.Errorf("config: tool_sets.%s: empty body", k)
		}
		ts.ExternalID = k
		for tk, tool := range ts.Tools {
			if tk == "" {
				return fmt.Errorf("config: tool_sets.%s.tools: empty external_id key", k)
			}
			if tool == nil {
				return fmt.Errorf("config: tool_sets.%s.tools.%s: empty body", k, tk)
			}
			tool.ExternalID = tk
		}
	}
	for k, ml := range cfg.MemoryLayers {
		if k == "" {
			return fmt.Errorf("config: memory_layers: empty external_id key")
		}
		if ml == nil {
			return fmt.Errorf("config: memory_layers.%s: empty body", k)
		}
		ml.ExternalID = k
		for ek, entry := range ml.Entries {
			if ek == "" {
				return fmt.Errorf("config: memory_layers.%s.entries: empty key", k)
			}
			if entry != nil {
				entry.Key = ek
			}
		}
	}
	for k, a := range cfg.Agents {
		if k == "" {
			return fmt.Errorf("config: agents: empty external_id key")
		}
		if a == nil {
			return fmt.Errorf("config: agents.%s: empty body", k)
		}
		a.ExternalID = k
		for vk, v := range a.Variations {
			if vk == "" {
				return fmt.Errorf("config: agents.%s.variations: empty external_id key", k)
			}
			if v == nil {
				return fmt.Errorf("config: agents.%s.variations.%s: empty body", k, vk)
			}
			v.ExternalID = vk
		}
	}
	return nil
}

// validateRefs checks intra-document constraints that can't be expressed in
// the YAML schema itself. Inter-document references (e.g. a variation
// assignment pointing at a tool set that neither exists in workspace nor in
// the config) are checked during Build after the plan lookups.
func validateRefs(cfg *Config) error {
	for _, a := range cfg.Agents {
		for _, v := range a.Variations {
			if v.Assignments == nil {
				continue
			}
			for i, ref := range *v.Assignments {
				count := 0
				if ref.Tool != "" {
					count++
				}
				if ref.ToolSet != "" {
					count++
				}
				if ref.Agent != "" {
					count++
				}
				if count == 0 {
					return fmt.Errorf("config: agents.%s.variations.%s.assignments[%d]: must set one of tool/tool_set/agent", a.ExternalID, v.ExternalID, i)
				}
				if count > 1 {
					return fmt.Errorf("config: agents.%s.variations.%s.assignments[%d]: must set exactly one of tool/tool_set/agent", a.ExternalID, v.ExternalID, i)
				}
			}
		}
	}
	return nil
}

func envFromOSEnviron() map[string]string {
	pairs := os.Environ()
	env := make(map[string]string, len(pairs))
	for _, p := range pairs {
		if i := strings.IndexByte(p, '='); i >= 0 {
			env[p[:i]] = p[i+1:]
		}
	}
	return env
}

