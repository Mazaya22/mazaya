package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("calculate a cone")
	fmt.Println("enter a side")
	conearea()
}

/* sloping height (t) fingers (j)  */

func conearea() {
	var t float64
	var r float64

	fmt.Println("input umber you want to count")
	fmt.Scan(&r)
	fmt.Scan(&t)

	slopingheight := math.Sqrt(t*t + r*r)
	base := math.Pi * r * r
	blanket := math.Pi * r * slopingheight
	cone := base + blanket

	fmt.Println("the area of the slant height is", slopingheight)
	fmt.Println("the area of the base area is", base)
	fmt.Println("the area of the blanket is", blanket)
	fmt.Println("the area of the cone is ", cone)

}
