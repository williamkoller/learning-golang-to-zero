package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("Hello")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel, i)
	}
}

func write(txt string) <-chan string {
	channel := make(chan string)

	go func () {
		for {
			channel <- fmt.Sprintf("value received: %s", txt)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel

}