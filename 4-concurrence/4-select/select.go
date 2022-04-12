package main

import (
	"fmt"
	"time"
)

func main() {
	channelOne, channelTwo := make(chan string), make(chan string)

	go func () {
		for {
			time.Sleep(time.Millisecond * 500)
			channelOne <- "channel one"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channelTwo <- "channel two"
		}
	}()

	for {
		select {
			case messageChannelOne := <- channelOne:
				fmt.Println(messageChannelOne)
			case messageChannelTwo := <- channelTwo:
				fmt.Println(messageChannelTwo)
		}
	}
}