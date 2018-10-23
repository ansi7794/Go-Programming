//Question
// For this exercise:
// 1. Create your own type. Have the underlying type be an int
// 2. Create a variable of new type with the identifier "x" using the "VAR"
// 3. in func main
// a. print out the value of the variable "x"
// b. print out the type of the variable "x"
// c. assign 42 to the variable "x" using the "=" operator
// d. print out the value of the variable "x"

package main

import "fmt"

type abc int

var x abc

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}