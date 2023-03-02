module modernc.org/sqlite

go 1.16

require (
	github.com/mattn/go-sqlite3 v1.14.10
	golang.org/x/sys v0.1.0
	golang.org/x/tools v0.2.0 // indirect
	modernc.org/ccgo/v3 v3.16.13
	modernc.org/libc v1.21.4
	modernc.org/mathutil v1.5.0
	modernc.org/tcl v1.10.0
	modernc.org/z v1.2.21
)

replace modernc.org/libc => ../libc

replace modernc.org/memory => ../memory
