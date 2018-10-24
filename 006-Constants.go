package main

import "fmt"

//constants without type consolidated
const (
	a = 22
	b float64 = 22.33 // with type
)
const c = "James Bond"

// constant - predeclared identifier
// resets with const keyword
const (
	x = iota
	y // event this is iota
	z
)

func main(){
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
