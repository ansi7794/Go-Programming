//Building on the code from the previous exercise
// 1. at the package level scope, using the 'var' keyword,
// create a variable with the identifier "y". The varible
// should be of the underlying type of your custom type of "x"
// 2. in func main
// a. this should already be done
// i. print out the value of the variable "x"
// ii. print out the type of the variable "x"
// iii. assign 42 to the variable "x" using the "=" operator
// iv. print out the value of the variable "x"
// b. now do this:
// i. now use conversion yo covert the type pf the value stored
// in x to the underlying type
// then assign that value to "y"
// print y and type of "y"

package main

import "fmt"

type abc int

var x abc
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
