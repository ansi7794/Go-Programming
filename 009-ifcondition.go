package main

import "fmt"

func main() {
 	if x := 42; x == 40 {
 		fmt.Println("x is 40")
	} else if x == 41 {
		fmt.Println("x is 41")
	} else {
		fmt.Println("x is 42")
	}
 }
