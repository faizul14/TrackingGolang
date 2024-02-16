package main

import (
	"fmt"
	// "strconv"
	"strings"
	"errors"
)

func validateInput(input string)(data bool, err error)  {

	defer fmt.Println("End") // defer alan di ekskusi pada kahir function

	if strings.TrimSpace(input) == ""{
		data = false
		err = errors.New("input cannot empty")
		return
	}

	return true, nil
}

func main() {
	var input string
	fmt.Print("Masukkan inputan int : ")
	fmt.Scanln(&input)

	var data bool
	var err error

	// data, err = strconv.Atoi(input)
		data, err = validateInput(input)


	if err == nil {
		fmt.Println("Data adalan", data)
	}else{
		panic(err.Error()) // pannic menmapilkan trace error dengan lengkap, dan kode setelahnya tidak akan di ekskusi kecuali deffer
		fmt.Println("Error")
	}
}