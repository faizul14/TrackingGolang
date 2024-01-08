package main

import (
	"fmt"
)

type dataDiri struct {
	name string
	age int
}

func main2() {
	var dataMap = map[string]interface{}{
		"name": "faezol",
		"age": 23,
	}

	fmt.Println(dataMap["name"])


	var varGeneral interface{}

	varGeneral = "faezol"
	fmt.Println(varGeneral)
	varGeneral = 14082000
	fmt.Println(varGeneral)

	varGeneral = varGeneral.(int) + 1
	fmt.Println(varGeneral)
}