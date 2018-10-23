package main

import "fmt"

//Global Scope - var keyword can be used to declare a variable outside of a function body
// declare + assign = initialize
var y = 33
//declare var with identifier z and type int
// assigns zero value
var z int

var a string

func main() {
	a = "This is an example problem"
	fmt.Println(a)
	//Declare a variable and assign a value - Short Declaration Operator
	//only works inside a function body
	x := 22
	fmt.Println(x)
	fmt.Println(y)
	ansi()
	fmt.Println(z)
}

func ansi() {
	fmt.Println("y inside ansi:", y)
}
