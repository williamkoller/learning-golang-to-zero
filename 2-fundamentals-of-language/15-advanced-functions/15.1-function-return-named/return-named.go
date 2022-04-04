package main

import "fmt"

func calcMaths(a, b int) (sum int, subtration int) {
	sum = a + b
	subtration = a - b
	return
}

func main() {
	sum, subtration := calcMaths(6, 8)
	fmt.Println(sum, subtration)
}