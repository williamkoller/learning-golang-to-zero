package main

import "fmt"

func sum(x, y int8) int8 {
	return x + y
}

func calcMath(x, y int8) (int8, int8) {
	total := x + y
	subtraction := x - y
	return total, subtraction
}

func main() {
	total := sum(4, 4)
	fmt.Println(total)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("text string")
	fmt.Println(result)

	// ignore second params
	resultSum, _ := calcMath(10, 15)
	fmt.Println("resultSum", resultSum)
	// fmt.Println("resultSubtraction", resultSubtraction)
}