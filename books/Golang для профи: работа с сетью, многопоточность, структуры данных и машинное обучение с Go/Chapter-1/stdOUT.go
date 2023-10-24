package main

import (
	"io"
	"os" // чтение командной строки
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1] + arguments[2]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")

}
