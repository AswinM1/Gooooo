package main

import "time"

// use the go keyword

func Hello() string {

	return "Hello"
}

func main() {

	var res string
	go Hello()
	time.Sleep(time.Second * 2)

}
