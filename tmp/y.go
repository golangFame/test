package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
)

// Example demonstrates the use of the trace package to trace
// the execution of a Go program. The trace output will be
// written to the file trace.out
func main() {

	fmt.Printf("go version %v\n", runtime.Version())
	f := os.Stdout

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %s", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %s", err)
	}
	defer trace.Stop()

	RunMyProgram()
}

func RunMyProgram() {
	fmt.Println("this function will be traced")
}
