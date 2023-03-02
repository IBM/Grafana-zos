// Copyright 2011 Evan Shaw. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE-MMAP-GO file.

//go:build darwin || dragonfly || freebsd || linux || openbsd || solaris || netbsd || zos
// +build darwin dragonfly freebsd linux openbsd solaris netbsd zos

// Modifications (c) 2017 The Memory Authors.

package memory // import "modernc.org/memory"

import (
	"os"
	"syscall"
	"unsafe"
)

const pageSizeLog = 12

var (
	osPageMask = osPageSize - 1
	osPageSize = 4096
	memMap     = make(map[*byte]string)
	dirname    = mkTempDir()
)

func mkTempDir() string {
	dname, err := os.MkdirTemp("", "mmap")
	if err != nil {
		panic("internal error")
	}
	return dname
}

func unmap(addr uintptr, size int) error {
	var s1 = struct {
		addr uintptr
		len  int
		cap  int
	}{addr, size, size}
	b := *(*[]byte)(unsafe.Pointer(&s1))

	if err := syscall.Munmap(b); err != nil {
		return err
	}

	filename, ok := memMap[(*byte)(unsafe.Pointer(addr))]
	if ok {
		delete(memMap, (*byte)(unsafe.Pointer(addr)))
		os.Remove(filename)
	}

	return nil
}

// pageSize aligned.
func mmap(size int) (uintptr, int, error) {
	size = roundup(size, osPageSize)

	// Create a temporary file to map the memory to
	f, err := os.CreateTemp(dirname, "allocated_mem")
	if err != nil {
		panic(err)
	}
	err = f.Truncate(int64(size + pageSize))
	if err != nil {
		panic(err)
	}
	fd := int(f.Fd())

	// golang.org/x/sys/unix doesn't have MAP_ANON, but the value defined in the IBM port is 0x0
	b, err := syscall.Mmap(fd, 0, size+pageSize, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		return 0, 0, err
	}
	memMap[&b[0]] = f.Name()

	n := len(b)
	p := uintptr(unsafe.Pointer(&b[0]))
	if p&uintptr(osPageMask) != 0 {
		panic("internal error")
	}

	mod := int(p) & pageMask
	if mod != 0 {
		m := pageSize - mod
		if err := unmap(p, m); err != nil {
			return 0, 0, err
		}

		n -= m
		p += uintptr(m)
	}

	if p&uintptr(pageMask) != 0 {
		panic("internal error")
	}

	// This was to make sure that we were not mapping more memory than we needed, but since
	// our page size is only 4KB we don't need to worry about it
	if n-(size+pageSize) != 0 {
		if err := unmap(p+uintptr(size+pageSize), n-(size+pageSize)); err != nil {
			return 0, 0, err
		}
	}

	return p, n, nil
}
