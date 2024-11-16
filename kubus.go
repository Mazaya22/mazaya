package main

import (
	"fmt"
	"math"
)

func luaskubus(sisi int) float64 {
	return 6 * math.Pow(5, 2)

}

func main() {
	var _ float64
	fmt.Println("masukan sisi kubus: ")
	var sisi = 5
	luas := luaskubus(sisi)

	fmt.Printf("luas permukaan kubus dengan sisi %.v adalah %.2f\n", sisi, luas)
}
