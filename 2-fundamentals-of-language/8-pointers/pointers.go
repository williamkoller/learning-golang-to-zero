package main

import "fmt"

func main() {
	var variableOne int = 10
	var variableTwo int = variableOne
	println(variableOne, variableTwo)

	variableOne++

	println(variableOne, variableTwo)

	// pointer is a memory reference
	var variableThree int
	var pointer *int

	variableThree = 100
	pointer = &variableThree

	fmt.Println(variableThree, pointer, *pointer)
	
	variableThree = 150
	fmt.Println(variableThree, pointer, *pointer)

}