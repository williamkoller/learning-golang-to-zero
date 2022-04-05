package main

import "fmt"

func funcOne() {
	fmt.Println("execute func one")
}

func funcTwo() {
	fmt.Println("execute func two")
}

func approvedStudent(a, b float32) bool {
	defer fmt.Println("calculated average. result will be returned")
	media := (a + b) / 2

	if media >= 6 {
		return  true
	}
	
	return false
}

func main() {
	// defer -> adiar
	defer funcOne()
	funcTwo()

	fmt.Println(approvedStudent(7, 8))
}