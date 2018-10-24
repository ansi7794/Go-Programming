// create and print a variable with raw string literal
// using iota create 4 constants for the last 5 years

package main

import "fmt"

const(
	m = iota + 2014
	n = iota + 2014
	o = iota + 2014
	p = iota + 2014
)

func main() {
	a := `THis is an "example" of
			raw string literal`
	fmt.Println(a)

	fmt.Println(m, n, o, p)
}