package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("execution recovered successfully")
	}
}

func approvedStudent(a, b float64) bool {
	defer recoverExecution()
	media := (a + b) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("the average is exactly six")
}

func main() {
	fmt.Println(approvedStudent(6, 6))
}