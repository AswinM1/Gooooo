package main

import "fmt"

func main() {
	channel := make(chan string)

	go func() {
		channel <- "hello from channel"

	}()
	data := <-channel
	fmt.Print(data)

}
