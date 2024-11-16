package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hitung luas krucut jika tinggi miringnya 4 dan jari jarinya 8")
	luaskrucut()
}

/* tinggi miring (t) = 4 jari jari (j) = 8 */

func luaskrucut() {
	var t float64
	var r float64

	t = 4
	r = 8

	tinggimiring := math.Sqrt(t*t + r*r)
	luasalasan := math.Pi * r * r
	luasselimut := math.Pi * r * tinggimiring
	luaskrucutan := luasalasan + luasselimut

	fmt.Println("tinggi miring nya adalah", tinggimiring)
	fmt.Println("luas alasnya adalah", luasalasan)
	fmt.Println("luas selimutnya adalah", luasselimut)
	fmt.Println("luas krucutnya adalah", luaskrucutan)

}
