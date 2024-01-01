package main

import "fmt"

/*
Deklarasi Variabel Menggunakan keyword new

-
*/

func main() {
	name := new(string)
	fmt.Println(name) // 0x20818a220
	fmt.Println(*name) // ""
}