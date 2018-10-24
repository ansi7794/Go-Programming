package main

import "fmt"

const(
	_ = iota
	KBo = 1 << (iota*10)
	MBo = 1 << (iota*10)
	GBo = 1 << (iota*10)
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	//shift left
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	//kilobyte and megabyte and gigabyte
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	// shifting with iota
	fmt.Printf("%d\t\t\t%b\n", KBo, KBo)
	fmt.Printf("%d\t\t\t%b\n", MBo, MBo)
	fmt.Printf("%d\t\t%b\n", GBo, GBo)
}