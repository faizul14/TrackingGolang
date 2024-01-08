package main

import "fmt"

type liveIn struct {
	adress string
	no int
}

type person struct {
	name string
	age int
	liveIn
}

func main()  {
	
	student := person{}

	student.name = "faezol padli"
	student.liveIn.adress = "Jl. Halmahera"
	student.no = 00


	fmt.Printf("Name: %v\nJalan: %v, no rumah %v",student.name, student.adress, student.liveIn.no)

}

