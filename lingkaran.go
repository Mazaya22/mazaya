package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("menghitung luas lingkaran")
	hitunglingkaran()
}

/*r 4*/
func hitunglingkaran() {
	var r float64
	r = 4
	luaslingkaran := math.Pi * r * r

	fmt.Println("luas lingkaran jika r = 4 adalah", luaslingkaran)
}
