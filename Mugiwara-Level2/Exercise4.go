// assigns an int to a variable
// prints that int in decimal, binary and hex
// shifts the bits of that int over 1 position
// prints that variable as decimal, binary and hex

package main

import "fmt"

func main(){
	var a int = 5
	b := a << 1

	fmt.Printf("a :: Decimal: %d Binary: %b Hexadecimal: %#x \n", a, a, a)
	fmt.Printf("b :: Decimal: %d Binary: %b Hexadecimal: %#x\n", b, b, b)
}