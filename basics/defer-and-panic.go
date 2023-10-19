package main

import "fmt"

func main() {
	defer finish()
	fmt.Println("Start program")
	fmt.Println("Process program")
	fmt.Println(deleter(5, 10))
	fmt.Println(deleter(5, 0))
}

func finish() {
	fmt.Println("Program finished")
}

func deleter(a, b int) int {
	if b == 0 {
		panic("Division ZERO")
	}
	return a / b
}
