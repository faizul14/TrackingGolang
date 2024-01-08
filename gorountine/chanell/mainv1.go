package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	// membuat chanell
	var message = make(chan string)

	var sayHelloTo = func (who string)  {
		var data = "Hello" + who

		message <- data

	}

	go sayHelloTo("Faezol")
	go sayHelloTo("Padli")
	go sayHelloTo("AnNis")

	var message1 = <- message
	fmt.Println("message1 runing")
	fmt.Println(message1)

	var message2 = <- message
	fmt.Println("message2 runing")
	fmt.Println(message2)

	var message3 = <- message
	fmt.Println("message3 runing")
	fmt.Println(message3)
}