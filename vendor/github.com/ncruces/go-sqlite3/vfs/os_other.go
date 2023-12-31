//go:build !linux && (!darwin || sqlite3_bsd)

package vfs

import (
	"io"
	"os"
)

func osSync(file *os.File, fullsync, dataonly bool) error {
	return file.Sync()
}

func osAllocate(file *os.File, size int64) error {
	off, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}
	if size <= off {
		return nil
	}
	return file.Truncate(size)
}
