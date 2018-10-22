package main

import "fmt"

func main() {
	fmt.Println("Hello Everyone!")
	foo()
	
	fmt.Println("Something More")
	
	for i := 0; i < 100; i++ {
		if i % 10 == 0 {
			fmt.Println(i);
		}
	}
}

func foo() {
	fmt.Println("I'm in foo")
}
