package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("THis wont print 1")
	case (2 == 2):
		fmt.Println("This is true 1")
		fallthrough //used to check all cases below
	case true:
		fmt.Println("This prints as well because of fallthrough")
		fallthrough
	case false:
		fmt.Println("This prints even though its false because of fallthrough")
	default:
		fmt.Println("None of them were true")
	}
}