package main

import "fmt"

func main() {
	//length of array is part of its type x and y are of different types
	var x [5]int
	var y [7]int
	fmt.Println(x)
	x[3] = 40
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(y)
	y[3] = 40
	fmt.Println(y)
	fmt.Println(len(y))
}
