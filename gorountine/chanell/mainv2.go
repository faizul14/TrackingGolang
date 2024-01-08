package main

import (
	"fmt"
	"runtime"
)

func printMessage(message chan string)  {
	fmt.Println(<- message)
}

func main() {
	runtime.GOMAXPROCS(2)

	// membuat chanell
	var message = make(chan string)

	for _, each := range []string{"Faezol", "Padli", "AnNis"} {
		go func(who string)  {
			var data = "Hello " + who
			message <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(message)
	}
}