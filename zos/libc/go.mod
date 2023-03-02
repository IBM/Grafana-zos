module modernc.org/libc

go 1.16

// replace modernc.org/libc => /home/joonl/dev/libc
// replace golang.org/x/sys => /home/joonl/go-work/src/golang.org/x/sys 

require (
	github.com/google/uuid v1.3.0
	github.com/mattn/go-isatty v0.0.12
	golang.org/x/sys v0.0.0-20211007075335-d3039528d8ac
	modernc.org/ccgo/v3 v3.12.47 // indirect
	modernc.org/mathutil v1.4.1
	modernc.org/memory v1.0.5
)

replace modernc.org/memory => ../memory
