package main

import (
	"fmt"
	"runtime"
	"time"
)

func print(n int, message string)  {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	go print(5, "hello")
	print(1000000, "apa kabar")

	// var input string
	// fmt.Scanln(&input)
	time.Sleep(1 * time.Second)

}