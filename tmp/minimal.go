package main

import (
	"fmt"
	"io"
	"runtime/trace"
)

func main() {
	trace.Start(io.Discard)
	fmt.Println("this will be traced")
	trace.Stop()
}