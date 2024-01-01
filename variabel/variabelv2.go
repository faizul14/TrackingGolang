package main

import "fmt"

/*
Deklarasi Variabel Tanpa Tipe Data

<nama-variabel> := <nilai>
*/

func main() {
	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	var firstName = "john"
	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	lastName := "wick"
	lastName = "ethan"
	lastName = "bourne"
	// := hanya sekali di gunakan ketika pembuatan pertama saja, selanjutnya ketika edit value gunakan =
	fmt.Printf("halo %s %s!\n", firstName, lastName)
}