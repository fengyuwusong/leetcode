package main

import (
	"fmt"
)

type A struct {
	AA string
}

type B struct {
	A *A
}

type point *A

func main() {
	b := B{A: nil}
	var p point
	fmt.Printf("%T\n", &b.A)
	fmt.Printf("%T\n", &p)
	fmt.Printf("%p\n", &b.A)
	newA := &A{
		AA: "123",
	}
	b.A = newA
	fmt.Printf("%p\n", newA)
	fmt.Printf("%p\n", b.A)
}
