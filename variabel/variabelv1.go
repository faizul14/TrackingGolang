package main

import "fmt"

/*
Deklarasi Variabel Beserta Tipe Data

var <nama-variabel> <tipe-data>
var <nama-variabel> <tipe-data> = <nilai>
*/

func main() {
	var firstName string = "john"
	var lastName string
	lastName = "wick"
	fmt.Printf("halo %s %s!\n", firstName, lastName)
}