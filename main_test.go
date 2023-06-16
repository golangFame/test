package main

import (
	"fmt"
	"testing"
)

const BYE = "rsetn"

func init(){
	fmt.Println("init")
}

func Test_HEY(t *testing.T){
	fmt.Println(BYE)
}