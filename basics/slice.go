package main

import "fmt"

func main() {
	var users []string
	users = append(users, "Nikita")
	fmt.Println(users)

	initialUsers := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users1 := initialUsers[1:4]
	users2 := initialUsers[4:6]
	users3 := initialUsers[6:]
	fmt.Println(users1)
	fmt.Println(users2)
	fmt.Println(users3)

	initialUsers = append(initialUsers[:2], initialUsers[3:]...)
	fmt.Println(initialUsers)
}
