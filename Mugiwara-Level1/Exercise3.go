// Question
// Using the code from previous exercise,
// 1. At the package level scope, assign the following
// values to the three variables
// a. for "x" assign 42
// b. for "y" assign "James Bond"
// c. for "z" assign true
// 2. in func main
// a. use fmt.Sprintf to print all of the values to one string.
// assign the returned value to type string using the short
// declaration operator to a variable with the identifier "s"

package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v",x,y,z)
	fmt.Println(s)
}
