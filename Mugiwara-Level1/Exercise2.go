// Question
// 1. Use var to declare three variables. The variables should
// have package level scope. Do not assign values to the variables.
// Use the following identifiers for the variables and make
// sure the variables store values of the following type
// a. identifier "x" type int
// b. identifier "y" type string
// c. identifier "z" type bool
// 2. in func main
// a. print out the values for each identifier
// b. the compiler assigned values to the variables.
// What are these values called?

package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("z =", z)

	// 2b. Zero values.
}
