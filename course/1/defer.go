package main

import "fmt"

func print() {
	defer fmt.Println("World")

	fmt.Println("hello")
}

func main() {
	print()
}
