package main

import (
	"errors"
	"fmt"
)

func main() {
	// int 8 - in bytes
	var number8 int8 = 127
	fmt.Println("int8", number8)

	// int 16 - in bytes
	var numberBytes int16 = 29999
	fmt.Println("int16", numberBytes)

	// int 32 - in bytes
	var number32 int32 = 1999999999
	fmt.Println("int32", number32)

	// int 64 - in bytes
	var number64 int64 = 8999999999999999999
	fmt.Println("int64", number64)

	// uint 8 - does not accept sign, for example negative numbers
	var numberUnit uint8 = 127
	fmt.Println("uint8", numberUnit)

	// alias
	// int 32 = rune
	var numberAlias rune = 1999999999
	fmt.Println("rune", numberAlias)

	// byte = unit8, does not accept sign, for example negative numbers
	var numberByte byte = 127
	fmt.Println("byte", numberByte)

	// float 32 - in bytes
	var numberFloat32 float32 = 9.99
	fmt.Println("float32", numberFloat32)

	// float 64 - in bytes
	var numberFloat64 float64 = 999999.999999999
	fmt.Println("float64", numberFloat64)

	// type inference, pick up according to machine architecture
	numberInferencia := 999999.99999999
	fmt.Println("type inference", numberInferencia)

	// string
	var str string = "Text"
	fmt.Println("string", str)

	// type inference string
	strTwo := "Text two"
	fmt.Println("type inference string", strTwo)

	// using rune -> int 32, char - number of table ascii
	char := 'B'
	fmt.Println("table ascii", char)

	// string with uninitialized value is empty
	var text string
	fmt.Println("string with uninitialized value is empty", text)

	// int with uninitialized value is empty
	var numberUninitialized int8
	fmt.Println("int with uninitialized value is 0", numberUninitialized)

	// bool = true
	var valueTrue bool = true
	fmt.Println("bool", valueTrue)

	// bool = false
	var valueFalse bool = false
	fmt.Println("bool", valueFalse)

	// bool with uninitialized value is empty = false
	var value bool
	fmt.Println("bool with uninitialized value is false", value)

	// nil = 0, error with uninitialized value is empty
	var erro error
	fmt.Println("error with uninitialized value is nil", erro)

	// custom error in go
	var errErrors error = errors.New("Errors Intern")
	fmt.Println(errErrors)
}
