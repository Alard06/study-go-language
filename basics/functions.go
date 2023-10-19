package main

import "fmt"

func main() {
	//xGenerated(5, 10)
	fmt.Println(sumAllElement(5, 10, 15, 20, 25, 30))

}

func sumAllElement(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func sum(a int, b int) (int, string) {
	return a + b, "integer"
}

func xGenerated(a int, b int) {
	for x := 0.0; x < 10; x += 0.5 {
		equation(x, a, b)
	}
}

func equation(x float64, a int, b int) {
	a2 := float64(a)
	b2 := float64(b)
	fmt.Println(a2*x*x + x*b2)
}
