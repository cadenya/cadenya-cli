//go:build unix

package main

import (
	"io/fs"
	"os"
	"syscall"
)

// checkPrivateFile requires the file to be owned by the current user with no
// group or other access. A git checkout writes 0644, so a committed or
// planted file fails here.
func checkPrivateFile(info fs.FileInfo) error {
	if perm := info.Mode().Perm(); perm&0o077 != 0 {
		return untrusted("its mode is %04o, not 0600", perm)
	}
	if st, ok := info.Sys().(*syscall.Stat_t); ok && int(st.Uid) != os.Getuid() {
		return untrusted("it is owned by another user")
	}
	return nil
}
