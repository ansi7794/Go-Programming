//Print the years you have been alive using for without a condition (for { } )

package main

import "fmt"

func main() {
	x := 1994
	for {
		fmt.Println(x)
		if x == 2018 {
			break
		}
		x++
	}
}