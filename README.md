go test -bench . -benchmem
go test -bench . -benchmem -gcflags=-l

The -l flag disables inlining
