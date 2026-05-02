// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestBulkWorkspaceResourcesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bulk-workspace-resources", "retrieve",
			"--id", "id",
		)
	})
}

func TestBulkWorkspaceResourcesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bulk-workspace-resources", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "0",
			"--managed-by-key", "managedByKey",
			"--sort-order", "sortOrder",
			"--state", "STATE_UNSPECIFIED",
		)
	})
}

func TestBulkWorkspaceResourcesApply(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bulk-workspace-resources", "apply",
			"--data", "{managedByKey: managedByKey, agents: {foo: {name: name, spec: {status: AGENT_STATUS_UNSPECIFIED, variationSelectionMode: VARIATION_SELECTION_MODE_UNSPECIFIED, description: description, inputDataSchema: {}, webhookEventsUrl: webhookEventsUrl}, labels: {foo: string}, schedules: {foo: {name: name, spec: {initialMessage: initialMessage, schedule: {calendars: [{comment: comment, dayOfMonth: [{end: 0, start: 0, step: 0}], dayOfWeek: [{end: 0, start: 0, step: 0}], hour: [{end: 0, start: 0, step: 0}], minute: [{end: 0, start: 0, step: 0}], month: [{end: 0, start: 0, step: 0}], second: [{end: 0, start: 0, step: 0}]}], intervals: [{every: '-160513s', offset: '-160513s'}], timezone: timezone}, data: {}, overlapPolicy: OVERLAP_POLICY_UNSPECIFIED, status: AGENT_SCHEDULE_STATUS_UNSPECIFIED, variationId: variationId}, labels: {foo: string}}}, variations: {foo: {name: name, spec: {compactionConfig: {summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}, constraints: {maxSubObjectives: 0, maxToolCalls: 0}, description: description, enableEpisodicMemory: true, episodicMemoryTtl: 0, modelConfig: {modelId: modelId, temperature: 0}, progressiveDiscovery: {hints: [string], maxTools: 0, rerankThreshold: 0}, prompt: prompt, weight: 0}, assignments: [{subAgentId: subAgentId, toolId: toolId, toolSetId: toolSetId}], labels: {foo: string}, memoryLayers: [{memoryLayerId: memoryLayerId, position: 0}]}}}}, memoryLayers: {foo: {name: name, spec: {type: MEMORY_LAYER_TYPE_UNSPECIFIED, description: description}, entries: {foo: {key: key, content: content, description: description, uploadId: uploadId}}, labels: {foo: string}}}, sourceUrl: sourceUrl, toolSets: {foo: {name: name, spec: {adapter: {http: {baseUrl: baseUrl, headers: {foo: string}}, mcp: {excludeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, headers: {foo: string}, includeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, toolApprovals: {always: true, only: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}}, url: url}}, description: description}, labels: {foo: string}, tools: {foo: {name: name, spec: {config: {http: {requestMethod: HTTP_METHOD_UNSPECIFIED, headers: {foo: string}, path: path, query: query, requestBodyContentType: requestBodyContentType, requestBodyTemplate: requestBodyTemplate, toolName: toolName}, mcp: {toolDescription: toolDescription, toolName: toolName, toolTitle: toolTitle}}, description: description, parameters: {foo: bar}, status: TOOL_STATUS_UNSPECIFIED, requiresApproval: true}, labels: {foo: string}}}}}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(bulkWorkspaceResourcesApply)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bulk-workspace-resources", "apply",
			"--data.managed-by-key", "managedByKey",
			"--data.agents", "{foo: {name: name, spec: {status: AGENT_STATUS_UNSPECIFIED, variationSelectionMode: VARIATION_SELECTION_MODE_UNSPECIFIED, description: description, inputDataSchema: {}, webhookEventsUrl: webhookEventsUrl}, labels: {foo: string}, schedules: {foo: {name: name, spec: {initialMessage: initialMessage, schedule: {calendars: [{comment: comment, dayOfMonth: [{end: 0, start: 0, step: 0}], dayOfWeek: [{end: 0, start: 0, step: 0}], hour: [{end: 0, start: 0, step: 0}], minute: [{end: 0, start: 0, step: 0}], month: [{end: 0, start: 0, step: 0}], second: [{end: 0, start: 0, step: 0}]}], intervals: [{every: '-160513s', offset: '-160513s'}], timezone: timezone}, data: {}, overlapPolicy: OVERLAP_POLICY_UNSPECIFIED, status: AGENT_SCHEDULE_STATUS_UNSPECIFIED, variationId: variationId}, labels: {foo: string}}}, variations: {foo: {name: name, spec: {compactionConfig: {summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}, constraints: {maxSubObjectives: 0, maxToolCalls: 0}, description: description, enableEpisodicMemory: true, episodicMemoryTtl: 0, modelConfig: {modelId: modelId, temperature: 0}, progressiveDiscovery: {hints: [string], maxTools: 0, rerankThreshold: 0}, prompt: prompt, weight: 0}, assignments: [{subAgentId: subAgentId, toolId: toolId, toolSetId: toolSetId}], labels: {foo: string}, memoryLayers: [{memoryLayerId: memoryLayerId, position: 0}]}}}}",
			"--data.memory-layers", "{foo: {name: name, spec: {type: MEMORY_LAYER_TYPE_UNSPECIFIED, description: description}, entries: {foo: {key: key, content: content, description: description, uploadId: uploadId}}, labels: {foo: string}}}",
			"--data.source-url", "sourceUrl",
			"--data.tool-sets", "{foo: {name: name, spec: {adapter: {http: {baseUrl: baseUrl, headers: {foo: string}}, mcp: {excludeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, headers: {foo: string}, includeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, toolApprovals: {always: true, only: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}}, url: url}}, description: description}, labels: {foo: string}, tools: {foo: {name: name, spec: {config: {http: {requestMethod: HTTP_METHOD_UNSPECIFIED, headers: {foo: string}, path: path, query: query, requestBodyContentType: requestBodyContentType, requestBodyTemplate: requestBodyTemplate, toolName: toolName}, mcp: {toolDescription: toolDescription, toolName: toolName, toolTitle: toolTitle}}, description: description, parameters: {foo: bar}, status: TOOL_STATUS_UNSPECIFIED, requiresApproval: true}, labels: {foo: string}}}}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"data:\n" +
			"  managedByKey: managedByKey\n" +
			"  agents:\n" +
			"    foo:\n" +
			"      name: name\n" +
			"      spec:\n" +
			"        status: AGENT_STATUS_UNSPECIFIED\n" +
			"        variationSelectionMode: VARIATION_SELECTION_MODE_UNSPECIFIED\n" +
			"        description: description\n" +
			"        inputDataSchema: {}\n" +
			"        webhookEventsUrl: webhookEventsUrl\n" +
			"      labels:\n" +
			"        foo: string\n" +
			"      schedules:\n" +
			"        foo:\n" +
			"          name: name\n" +
			"          spec:\n" +
			"            initialMessage: initialMessage\n" +
			"            schedule:\n" +
			"              calendars:\n" +
			"                - comment: comment\n" +
			"                  dayOfMonth:\n" +
			"                    - end: 0\n" +
			"                      start: 0\n" +
			"                      step: 0\n" +
			"                  dayOfWeek:\n" +
			"                    - end: 0\n" +
			"                      start: 0\n" +
			"                      step: 0\n" +
			"                  hour:\n" +
			"                    - end: 0\n" +
			"                      start: 0\n" +
			"                      step: 0\n" +
			"                  minute:\n" +
			"                    - end: 0\n" +
			"                      start: 0\n" +
			"                      step: 0\n" +
			"                  month:\n" +
			"                    - end: 0\n" +
			"                      start: 0\n" +
			"                      step: 0\n" +
			"                  second:\n" +
			"                    - end: 0\n" +
			"                      start: 0\n" +
			"                      step: 0\n" +
			"              intervals:\n" +
			"                - every: '-160513s'\n" +
			"                  offset: '-160513s'\n" +
			"              timezone: timezone\n" +
			"            data: {}\n" +
			"            overlapPolicy: OVERLAP_POLICY_UNSPECIFIED\n" +
			"            status: AGENT_SCHEDULE_STATUS_UNSPECIFIED\n" +
			"            variationId: variationId\n" +
			"          labels:\n" +
			"            foo: string\n" +
			"      variations:\n" +
			"        foo:\n" +
			"          name: name\n" +
			"          spec:\n" +
			"            compactionConfig:\n" +
			"              summarization:\n" +
			"                instructions: instructions\n" +
			"              toolResultClearing:\n" +
			"                preserveRecentResults: 0\n" +
			"              triggerThreshold: 0\n" +
			"            constraints:\n" +
			"              maxSubObjectives: 0\n" +
			"              maxToolCalls: 0\n" +
			"            description: description\n" +
			"            enableEpisodicMemory: true\n" +
			"            episodicMemoryTtl: 0\n" +
			"            modelConfig:\n" +
			"              modelId: modelId\n" +
			"              temperature: 0\n" +
			"            progressiveDiscovery:\n" +
			"              hints:\n" +
			"                - string\n" +
			"              maxTools: 0\n" +
			"              rerankThreshold: 0\n" +
			"            prompt: prompt\n" +
			"            weight: 0\n" +
			"          assignments:\n" +
			"            - subAgentId: subAgentId\n" +
			"              toolId: toolId\n" +
			"              toolSetId: toolSetId\n" +
			"          labels:\n" +
			"            foo: string\n" +
			"          memoryLayers:\n" +
			"            - memoryLayerId: memoryLayerId\n" +
			"              position: 0\n" +
			"  memoryLayers:\n" +
			"    foo:\n" +
			"      name: name\n" +
			"      spec:\n" +
			"        type: MEMORY_LAYER_TYPE_UNSPECIFIED\n" +
			"        description: description\n" +
			"      entries:\n" +
			"        foo:\n" +
			"          key: key\n" +
			"          content: content\n" +
			"          description: description\n" +
			"          uploadId: uploadId\n" +
			"      labels:\n" +
			"        foo: string\n" +
			"  sourceUrl: sourceUrl\n" +
			"  toolSets:\n" +
			"    foo:\n" +
			"      name: name\n" +
			"      spec:\n" +
			"        adapter:\n" +
			"          http:\n" +
			"            baseUrl: baseUrl\n" +
			"            headers:\n" +
			"              foo: string\n" +
			"          mcp:\n" +
			"            excludeTools:\n" +
			"              operator: OPERATOR_UNSPECIFIED\n" +
			"              filters:\n" +
			"                - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"                  matcher:\n" +
			"                    caseSensitive: true\n" +
			"                    contains: contains\n" +
			"                    endsWith: endsWith\n" +
			"                    exact: exact\n" +
			"                    regex: regex\n" +
			"                    startsWith: startsWith\n" +
			"            headers:\n" +
			"              foo: string\n" +
			"            includeTools:\n" +
			"              operator: OPERATOR_UNSPECIFIED\n" +
			"              filters:\n" +
			"                - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"                  matcher:\n" +
			"                    caseSensitive: true\n" +
			"                    contains: contains\n" +
			"                    endsWith: endsWith\n" +
			"                    exact: exact\n" +
			"                    regex: regex\n" +
			"                    startsWith: startsWith\n" +
			"            toolApprovals:\n" +
			"              always: true\n" +
			"              only:\n" +
			"                operator: OPERATOR_UNSPECIFIED\n" +
			"                filters:\n" +
			"                  - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"                    matcher:\n" +
			"                      caseSensitive: true\n" +
			"                      contains: contains\n" +
			"                      endsWith: endsWith\n" +
			"                      exact: exact\n" +
			"                      regex: regex\n" +
			"                      startsWith: startsWith\n" +
			"            url: url\n" +
			"        description: description\n" +
			"      labels:\n" +
			"        foo: string\n" +
			"      tools:\n" +
			"        foo:\n" +
			"          name: name\n" +
			"          spec:\n" +
			"            config:\n" +
			"              http:\n" +
			"                requestMethod: HTTP_METHOD_UNSPECIFIED\n" +
			"                headers:\n" +
			"                  foo: string\n" +
			"                path: path\n" +
			"                query: query\n" +
			"                requestBodyContentType: requestBodyContentType\n" +
			"                requestBodyTemplate: requestBodyTemplate\n" +
			"                toolName: toolName\n" +
			"              mcp:\n" +
			"                toolDescription: toolDescription\n" +
			"                toolName: toolName\n" +
			"                toolTitle: toolTitle\n" +
			"            description: description\n" +
			"            parameters:\n" +
			"              foo: bar\n" +
			"            status: TOOL_STATUS_UNSPECIFIED\n" +
			"            requiresApproval: true\n" +
			"          labels:\n" +
			"            foo: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"bulk-workspace-resources", "apply",
		)
	})
}
