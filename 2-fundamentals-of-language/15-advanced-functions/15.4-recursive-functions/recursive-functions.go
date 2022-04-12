package main

import "fmt"

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position - 2) + fibonacci(position - 1)
}

func main() {
	position := uint(12)

	for i := uint(1); i < position; i++ {
		fmt.Println(fibonacci(i))
	}
}