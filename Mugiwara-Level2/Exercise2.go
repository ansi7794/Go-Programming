//use operators ==, <=, >=, !=, <, >
package main

import "fmt"

func main(){
	a := (22 == 22)
	b := (22 <= 21)
	c := (22 >= 25)
	d := (22 != 24)
	e := (44 < 41)
	f := (44 > 41)

	fmt.Println(a, b, c, d, e, f)
}