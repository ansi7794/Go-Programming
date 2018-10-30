package main

import "fmt"

func main() {
	x := make([]int, 10, 12) //make(type length capacity)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))//print capacity

	x = append(x, 11)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 13)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) // beyond original capacity -- doubles capacity
}