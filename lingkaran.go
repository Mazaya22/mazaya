package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("calculate area of a circle")
	countcircles()
}

/*r 4*/
func countcircles() {
	var r float64

	fmt.Println("enter a radius you want:")
	fmt.Scan(&r)
	surfaceofcircle := math.Pi * r * r

	fmt.Println("circle area with radius", r, "is:", surfaceofcircle)
}
