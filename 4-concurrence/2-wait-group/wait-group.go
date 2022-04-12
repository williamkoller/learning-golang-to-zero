package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var awaitGroup sync.WaitGroup

	awaitGroup.Add(4)

	go func() {
		write("Hello")
		awaitGroup.Done()
	}()

	go func() {
		write("programming in go")
		awaitGroup.Done()
	}()

	go func() {
		write("goroutines three")
		awaitGroup.Done()
	}()

	go func() {
		write("goroutines four")
		awaitGroup.Done()
	}()

	awaitGroup.Wait()

}

func write(txt string) {
	for i := 0; i < 5; i++ {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}