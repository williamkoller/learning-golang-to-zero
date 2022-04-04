package main

import "fmt"

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func write(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	sum := sum(1, 2, 3, 4, 5, 6)

	write("redragon", 12, 34, 55, 99)
	fmt.Println(sum)
}