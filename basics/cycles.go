package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(i, i*i)
	}

	fmt.Println()

	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	users := [...]string{"Tom", "Nikita", "John"}

	for key, value := range users {
		fmt.Println(key, value)
	}
	fmt.Println(users)
	sum := 0
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue
		} else if i == 11 {
			break
		}
		sum += i
		fmt.Println(i, sum)
	}

}
