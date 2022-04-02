package main

import "fmt"

func main() {
	number := -3
	fifteen := 15

	if number > fifteen {
		fmt.Println("greater than fifteen")
	} else {
		fmt.Println("less than fifteen")
	}

	// if init, is most used to define a scope

	if otherNumber := number; otherNumber > 0 {
		fmt.Println("Number greater than zero")
	} else if number < -10 {
		fmt.Println("Number less than -10")
	} else {
		fmt.Println("Between 0 and -10")
	}
}