package main

import (
	"fmt"
	"time"
)

func write(txt string, channel chan string) {
	for i := 0; i< 5; i++ {
		channel <- txt
		time.Sleep(time.Second)
	}

	close(channel)
}

func main() {
	channel := make(chan string)

	go write("Hello World!", channel)

	fmt.Println("after the write function starts executing")

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("end of program")
}