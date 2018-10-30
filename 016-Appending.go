package main

import "fmt"

func main() {
	x := []int{40,15,33,67,98,55}
	fmt.Println(len(x))
	fmt.Println(x)
	x = append(x, 2, 4, 5, 7)
	fmt.Println(x)
}