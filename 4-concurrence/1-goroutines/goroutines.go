package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Hello World")
	write("Programming in Go")
}

func write(txt string) {
	for {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}