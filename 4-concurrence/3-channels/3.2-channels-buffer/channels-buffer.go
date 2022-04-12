package main

import "fmt"

func main() {
	// with 2 = buffer, if block with capacity all
	channel := make(chan string, 2)

	channel <- "Hello"
	channel <- "World"

	message := <- channel
	messageTwo := <- channel

	fmt.Println(message)
	fmt.Println(messageTwo)
}