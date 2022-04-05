package main

import "fmt"

// paramters for copy
func invertSignal(number int) int {
	return number * -1
}

// paramters for pointers
func invertSignalWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 20
	numberInvert := invertSignal(number)

	fmt.Println(numberInvert)
	fmt.Println(number)

	newNumber := 40
	fmt.Println(newNumber)

	invertSignalWithPointer(&newNumber)
	fmt.Println(newNumber)

}