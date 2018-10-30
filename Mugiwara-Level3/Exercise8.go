package main

import "fmt"

func main() {
	favSport := "football"
	switch favSport {
	case "cricket":
		fmt.Println("favSport is cricket")
	case "soccer":
		fmt.Println("favSport is soccer")
	case "football":
		fmt.Println("favSport is football")
	}
}
