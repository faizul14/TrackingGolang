package main

import "fmt"

/*
Deklarasi Multi Variabel

Go mendukung metode deklarasi banyak variabel secara bersamaan, caranya
dengan menuliskan variabel-variabel-nya dengan pembatas tanda koma (,).
Untuk pengisian nilainya-pun diperbolehkan secara bersamaan.
*/

func main() {
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	var fourth, fifth, sixth string = "empat", "lima", "enam"

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	fmt.Println(first, second, third, fourth, fifth, sixth, seventh, eight, ninth)
}