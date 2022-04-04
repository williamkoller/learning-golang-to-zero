package main

import "fmt"

func main() {
	result := func(txt string) string {
		return fmt.Sprintf("Received -> %s", txt)
	}("anonymous functions")

	fmt.Println(result)
}