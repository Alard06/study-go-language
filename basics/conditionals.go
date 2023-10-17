package main

import "fmt"

func main() {
	a := 10
	b := 10

	if a < b {
		fmt.Println("a < b")
	} else if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a == b")
	}

	x := 5

	switch x {
	case 5:
		fmt.Println("x == 5")
	case 6:
		fmt.Println("x == 6")
	}

	fmt.Println()
}
