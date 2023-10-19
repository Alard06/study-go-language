package main

import "fmt"

func main() {
	//var numbers [5]int = [5]int{1, 2}
	//fmt.Println(numbers)

	numbers := [5]int{2, 4, 5, 6}
	fmt.Println(numbers)
	numbers2 := [...]int{5, 5, 6, 7, 2, 1, 5}
	fmt.Println(len(numbers2))
	fmt.Println()

	numbers3 := [...]string{1: "I", 2: "II", 3: "III", 0: "0"}
	fmt.Println(numbers3[0])
	fmt.Println(numbers3[1])
	fmt.Println(numbers3[2])
	fmt.Println(numbers3[3])
}
