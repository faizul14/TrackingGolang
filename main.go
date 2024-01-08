package main

// import "fmt"

// import "GOLANGV1/library"
import (
	. "GOLANGV1/library" //import dengan prefix (.) = bisa langsung panggil tanpa menuliskan package nya
)


func main() {
	// var firstNmae string = "faezol"
	// lastName := "padli"

	// fmt.Println(firstNmae + lastName)

	// fmt.Printf("Nama depan saya %s dan nama belakang saya %s \n", firstNmae, lastName)

	// satu, dua, tiga := 1,2,3
	// fmt.Println(satu+dua,tiga)

	// name := new (int)
	// fmt.Println(name)

	// // seleksi kondisi
	// isSucces := true
	// fmt.Println(isSucces)
	SayHello()
	// library.introduce() //ini tidak bisa karna unexported
}