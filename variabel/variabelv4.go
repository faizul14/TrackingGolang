package main

import "fmt"

/*
Deklarasi Multi Variabel

Underscore (_) adalah reserved variable yang bisa dimanfaatkan untuk
menampung nilai yang tidak dipakai. Bisa dibilang variabel ini merupakan
keranjang sampah.
*/

func main() {
	_ = "belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "john", "wick"
}