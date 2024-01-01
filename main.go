package main

import "fmt"

func main() {
	var firstNmae string = "faezol"
	lastName := "padli"

	fmt.Println(firstNmae + lastName)

	fmt.Printf("Nama depan saya %s dan nama belakang saya %s \n", firstNmae, lastName)

	satu, dua, tiga := 1,2,3
	fmt.Println(satu+dua,tiga)

	name := new (int)
	fmt.Println(name)

	// seleksi kondisi
	isSucces := true
	fmt.Println(isSucces)
}