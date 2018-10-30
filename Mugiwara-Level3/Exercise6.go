// any program with if statement in action

package main

import "fmt"

func main() {
	for i:=0; i<=55; i++ {
		if i % 2 == 0 && i % 3 == 0 {
			fmt.Printf("%d is divisible by both 2 and 3\n", i)
		} else if i % 2 == 0 {
			fmt.Printf("%d is divisible by only 2\n", i)
		} else if i % 3 == 0 {
			fmt.Printf("%d is divisible by only 3\n", i)
		} else {
			fmt.Printf("%d is not divisible by 2 or 3\n", i)
		}
	}
}