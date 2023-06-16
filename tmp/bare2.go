package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	if err := runtime.StartTrace(); err != nil {
		log.Fatal("failed to start trace")
	}

	fmt.Println("this will be traced")
	runtime.StopTrace()
	fmt.Println("doesn't panic")
}
