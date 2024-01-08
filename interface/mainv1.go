package main

import (
	f "fmt"
	matematic "math"
)

type hitung interface {
	luas()	float64
	keliling() float64
}


// menghitung lingkara

type lingkaran struct {
	diameter float64
}


func (d lingkaran) jarijari() float64  {
	return d.diameter / 2
}

func (d lingkaran) luas() float64  {
	return matematic.Pi * matematic.Pow(d.jarijari(), 2)
}

func (d lingkaran) keliling() float64  {
	return matematic.Pi * d.diameter
}

// untuk menghitung persegi

type persegi struct {
	sisi float64
}

func (s persegi) luas() float64  {
	return matematic.Pow(s.sisi, 2)
}

func (s persegi) keliling() float64  {
	return 4 * s.sisi
}

func main() {
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	f.Println("===== Luas persegi =====")
	f.Println("Luas		:", bangunDatar.luas())
	f.Println("Keliling	:", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	f.Println("===== Luas Lingkaran =====")
	f.Println("Luas		:", bangunDatar.luas())
	f.Println("Keliling	:", bangunDatar.keliling())
	f.Println("Jari-Jari	:", bangunDatar.(lingkaran).jarijari())
}