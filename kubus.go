package main

import (
	"fmt"
)

func areaofthecube(side int) int {
	return 6 * side * side

}

func main() {
	var _ float64
	fmt.Println("enter the side of cubes: ")
	var side = 5
	surface := areaofthecube(side)

	fmt.Printf("surface area of a cube with side %.v is %.2f\n", side, surface)
}
