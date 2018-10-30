package main

import "fmt"

func main() {
	x := []int{40,15,33,67,98,55}
	fmt.Println(len(x))
	fmt.Println(x[1:])
	fmt.Println(x[1:3]) // from including -- up to not including
	fmt.Println()
}
