package main

import "fmt"

type myString rune

func main(){
	x:= myString(0x1A34B7)
	fmt.Printf("%q\n",x)
}
