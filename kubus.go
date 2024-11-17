package main

import (
	"fmt"
)

func main() {
	var customside float64

	fmt.Println("enter the side of cubes: ")

	fmt.Scan(&customside)

	customsidesurface := 6 * customside * customside

	fmt.Println("surface area of a cube with side", customside, "is:", customsidesurface)

	/* if you want to know surface of cube just input a number you want*/
}
