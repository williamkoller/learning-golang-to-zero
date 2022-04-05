package main

import "fmt"

var n int

func init() {
	fmt.Println("execute func init")
	n = 10
}

func main() {
	fmt.Println("execute func main")
	fmt.Println(n)
}