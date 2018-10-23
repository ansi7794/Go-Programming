package main

import (
	"fmt"
	"runtime"
)

//bool type
var x bool
// numeric type
var a int
var b float64
// string type
var s string

func main() {
	// bool type
	fmt.Println(x)
	x = true
	fmt.Println(x)
	//numeric type
	a = 42
	b = 42.32456
	fmt.Printf("%T ", a)
	fmt.Println(a)
	fmt.Printf("%T ", b)
	fmt.Println(b)

	s = "Hello"
	fmt.Println(s)
	s = `"Hello 
		how are you?"` //back ticks include the escape characters too
	fmt.Println(s)
	fmt.Printf("%T ", s)

	//byte representation of string
	s = "Hello! How are you?"
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	//utf8 character
	for i:= 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")
	for i, v := range s {
		fmt.Println(i, v)
		fmt.Printf("at index position %d we have hex %x \n", i, v)
	}

	// prints the runtime OS and Architecture
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}