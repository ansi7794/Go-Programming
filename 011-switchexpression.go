package main

import (
	"fmt"
)

func main() {
	c := 5

	switch c {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 0:
		fmt.Println("This is a numeric")
	default:
		fmt.Println("this is not a numeric")

	}
}