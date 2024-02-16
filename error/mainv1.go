package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Masukkan inputan int : ")
	fmt.Scanln(&input)

	var data int
	var err error

	data, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println("Data adalan", data)
	}else{
		fmt.Println("Error", err.Error())
	}
}