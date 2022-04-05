package main

import "fmt"

func closure() func() {
	txt := "inside closure function"

	function := func() {
		fmt.Println(txt)
	}

	return function
}

func main() {
	txt := "inside main function"
	fmt.Println(txt)

	newFunc := closure()

	newFunc()
}