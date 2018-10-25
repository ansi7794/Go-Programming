package main

import "fmt"

func main() {
	//for init; condition; post {}
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//nested loops
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Println(i, j)
		}
	}

	//for boolean condition
	x := 1
	for x < 10 {
		fmt.Println(x)
		x ++
	}

	// for with if and break and continue
	x = 1
	for {
		x++

		if x > 50 {
			break
		}

		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}