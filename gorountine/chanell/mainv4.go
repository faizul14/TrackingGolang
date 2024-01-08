package main

import (
	"fmt"
)

func sendMessage(ch chan<- string)  {

	for i := 0; i < 20; i++ {
		ch  <- fmt.Sprintf("data %d", i)
	}

	close(ch)
	
}

func printMessage(ch <-chan string)  {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	var message = make(chan string)
	go sendMessage(message)
	printMessage(message)
}