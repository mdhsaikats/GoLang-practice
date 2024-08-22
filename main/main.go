package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	var saikat string = "Saikat"
	var saikat2 = "Jane"
	x := 2

	fmt.Println(saikat)
	fmt.Println(saikat2)
	fmt.Println(x)

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var d, e, f int = 1, 2, 3

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	var g, h = 6, "Hello"
	i, j := 7, "World!"

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)

	var (
		ab int
		bc int    = 1
		cd string = "hello"
	)
	fmt.Println(ab)
	fmt.Println(bc)
	fmt.Println(cd)

	const PI = 3.14

	fmt.Println(PI)

	const PIx int = 1

	fmt.Println(PIx)

	const (
		A int = 1
		B     = 3.14
		C     = "Hi!"
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

}
