package main

import (
	"fmt"
)

func main() {

	fmt.Println("hallo world")
	menghitungLuasKotak()

}

/*
*
pr pertama, type data, int, string, float, double
pr2, buat 3 function

	1.hitung kubus
	2.itung lingkaran
	3.hitung krucut

..syarat tiap penghitungan menggunakan file yg terpisah , sehingga menjadi k.go, l.go, k.go
*/
func menghitungLuasKotak() {
	var x int
	x = 10
	y := 5
	luas := x * y

	var x2 = 10
	y2 := 3
	luas2 := x2 * y2

	var x3 = 3
	y3 := 15
	luas3 := x3 * y3
	fmt.Println("luas", luas)
	fmt.Println("luas", luas2)
	fmt.Println("luas", luas3)

}
