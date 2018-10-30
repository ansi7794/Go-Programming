package main

import "fmt"

func main() {
	x := []int{40,15,33,67,98,55}
	fmt.Println(len(x))
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}
}
