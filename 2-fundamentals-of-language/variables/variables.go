package main

import "fmt"

func main() {
	// com inferencia de tipo
	var variable_one string = "variable one"

	// sem inferencia de tipo
	variable_two := "variable two"

	fmt.Println(variable_one, variable_two)

	var (
		variable_three string = "variable three"
		variable_four string = "variable four"
	)

	fmt.Println(variable_three, variable_four)

	variable_five, variable_six := "variable five", "variable six"

	fmt.Println(variable_five, variable_six)

	const constant string = "constant"
	
	fmt.Println(constant)

	// trocar de valores
	variable_five, variable_six = variable_six, variable_five

	fmt.Println(variable_five, variable_six)

}
