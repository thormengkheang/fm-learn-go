package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg)
		time.Sleep(400 * time.Millisecond)
	}
	channel <- "Done"
}

func main() {
	var channel chan string
	// go printMessage("Goroutine is awesome!")
	go printMessage("Frontend masters is awesome!")
	// go printMessage("Yes, it is!")
	// time.Sleep(time.Minute)
	res := <-channel
	fmt.Println(res)
}
