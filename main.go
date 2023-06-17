package main

import "fmt"

type CommonOption[T any, C TypeA[T] | TypeB[T]] struct {
	value     T
	container *C
}

func (o *CommonOption[T, C]) WithValue(v T) *C {
	o.value = v

	return o.container
}

type TypeA[T any] struct {
	*CommonOption[T, TypeA[T]]

	subFieldA string
}

type TypeB[T any] struct {
	*CommonOption[T, TypeB[T]]
}

func main(){
	fmt.Println("Starting main")
}