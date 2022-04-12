package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexer(write("Hello"), write("Program go"))
	 for i := 0; i < 10; i++ {
		 fmt.Println(<-channel, i)
	 }
}

func multiplexer(channelInputOne, channelInputTwo <-chan string) <-chan string {
	channelOutput := make(chan string)

	go func () {
		for {
			select {
				case message := <- channelInputOne:
					channelOutput <- message
				case message := <- channelInputTwo:
					channelOutput <- message
			}
		}
	}()

	return channelOutput
}

func write(txt string) <-chan string {
	channel := make(chan string)

	go func () {
		for {
			channel <- fmt.Sprintf("value received: %s", txt)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel

}