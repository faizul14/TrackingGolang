package main

import (
	"fmt"
	"time"
	"os"
)

func timer(timer int, ch chan<- bool)  {
	time.Sleep(5 * time.Second)
	ch <- true
}

func watcher(timeOut int, ch <-chan bool)  {
	<- ch // blocking di sini
	fmt.Println("\n time out! no answer more than", timeOut, "second")
	os.Exit(0)
}

func main() {
	timeOut := 5
	var ch = make(chan bool)

	go timer(timeOut, ch)
	go watcher(timeOut, ch)

	var input string
	fmt.Print("what is 725/25 ? ")
	fmt.Scan(&input)
	if input == "29" {
		fmt.Println("the answer is right!")
	} else {
		fmt.Println("the answer is wrong!")
	}
}