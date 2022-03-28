package main

import "fmt"

func main() {
	// arithmetic (+, -, /, *, %)
	sum := 2 + 2
	subtraction := 6 - 4
	division := 30 / 2
	multiplication := 5 * 20
	restToDivision := 8 % 3

	fmt.Println(sum, subtraction, division, multiplication, restToDivision)

	// this is not allowed in go
	/**
	var numberOne int16 = 10
	var numberTwo int32 = 25
	total := numberOne + numberTwo
	fmt.Println(total)
	*/

	var numberOne int16 = 10
	var numberTwo int16 = 25
	var total int16 = numberOne + numberTwo
	fmt.Println(total)

	// assignment operator
	var valueAssignment string = "string"
	fmt.Println(valueAssignment)

	assignment := "assignment"
	fmt.Println(assignment)

	// relational operators
	fmt.Println(4 > 2)
	fmt.Println(4 >= 2)
	fmt.Println(3 < 2)
	fmt.Println(3 <= 2)
	fmt.Println(2 == 2)
	fmt.Println(2 != 2)

	// logical operators
	fmt.Println("---------")
	truth, falsy := true, false
	fmt.Println(truth && falsy)
	fmt.Println(truth || falsy)
	fmt.Println(!truth)
	fmt.Println(!falsy)

	// unary operators
	fmt.Println("---------")
	number := 10

	number++ // number = number + 1
	fmt.Println(number)

	number += 20 // number = number + 20
	fmt.Println(number)

	number--
	fmt.Println(number)

	number -= 20
	fmt.Println(number)

	number *= 40 // number = number * 40
	fmt.Println(number)

	number /= 20 // number = number / 20
	fmt.Println(number)

	number %= 8 // number = number % 8
	fmt.Println(number)

	// ternary operators

	/**
	this is not allowed in go

	txt := number > 5 ? "maior que 5" : "menor que 5"
	fmt.Println(txt)
	*/
	
	var txt string
	if number > 5 {
		txt = "maior que 5"
	} else {
		txt = "menor que 5"
	}

	fmt.Println(txt)

}
