package main

import (
	"fmt"

	_ "github.com/cilium/ebpf"
)

//go:generate env --unset=GOROOT go run github.com/cilium/ebpf/cmd/bpf2go 

func main(){
	fmt.Println("Starting main")
}