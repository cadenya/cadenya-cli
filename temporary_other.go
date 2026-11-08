//go:build !unix

package main

import "io/fs"

// checkPrivateFile has no portable owner or mode check off Unix.
func checkPrivateFile(fs.FileInfo) error { return nil }
