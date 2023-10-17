package main

// Переменные

import "fmt"

func main() {
	const pi = 3.14
	const e = 2.7

	const e2 = e

	const (
		a = 1
		c
		d
		r = 4
		key
	)

	fmt.Println(pi, e)
	var hello string
	hello = "hello world!"
	fmt.Println(hello)

	var world string = "World!"

	fmt.Println(world)

	var (
		age  int    = 19
		name string = "Nikita"
	)
	fmt.Println(age, name)

	hello = "OMG"
	fmt.Println(hello)

	// Краткое определение переменной

	myName := "Nikita"

	fmt.Println(myName)

}
