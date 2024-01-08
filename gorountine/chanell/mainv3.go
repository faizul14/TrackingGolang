package main

import (
	"fmt"
	"runtime"
)

func Average(data []int, ch chan float64)  {
	sum := 0

	for _, v := range data {
		sum += v
	}

	ch <- float64(sum) / float64(len(data))
}

func ValueMax(data []int, ch chan float64)  {
	max := data[0]

	for _, v := range data {
		if max < v {
			max = v
		}
	}

	ch <- float64(max)
}

func main() {

	runtime.GOMAXPROCS(2)

	data := []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println(data)

	var ch1 = make(chan float64)
	go Average(data, ch1)

	var ch2 = make(chan float64)
	go ValueMax(data, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <- ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <- ch2:
			fmt.Printf("Max \t: %.2f \n", max)
		}
		
	}
}