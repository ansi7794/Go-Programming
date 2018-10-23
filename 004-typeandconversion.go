package main

import "fmt"

var a int
type bloom int
var b bloom

func main() {
	a = 23
	b = 225

	fmt.Println(a)
	fmt.Printf("%T\n", a) //Print type with %T

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// a = b is not possible
	// as a is of type int and b is of type bloom
	// static type for any variable in Go
	// type conversion -- not type casting
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = bloom(a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
