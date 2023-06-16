package main

import (
	"fmt"
	"io"
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	
	if err := runtime.StartTrace(); err != nil {
		log.Fatal("failed to start trace")
	}
	go func() {
		w := io.Discard
		for {
			fmt.Println("running")
			data := runtime.ReadTrace()
			if data == nil {
				break
			}
			// fmt.Println(string(data))
			w.Write(data)
		}
	}()
	fmt.Println("this will be traced")

	runtime.StopTrace()
}
