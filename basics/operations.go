package main

import "fmt"

func main() {
	var a = 5
	var b = 10

	var c = a + b
	fmt.Println(c)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a--
	b--
	fmt.Println(a)
	fmt.Println(b)
	a++
	b++
	fmt.Println(b)
	fmt.Println(b)
}
