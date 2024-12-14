package main

import (
	"fmt"
	"math"
)

func countcube(s int) int {
	return 6 * s * s
}

func countcircle(r float64) float64 {
	return 3.14 * r * r
}

func countcone(hg, rd float64) float64 {
	ls := math.Sqrt(rd*rd + hg*hg)
	return math.Pi * rd * (rd + ls)
}

func main() {
	var choice int
	fmt.Println("select a geometric shape to calculate area: ")
	fmt.Println("1. surface of cube")
	fmt.Println("2. surface of circle")
	fmt.Println("3. surface of cone")
	fmt.Print("enter options (1,2,3): ")
	fmt.Scan(&choice)

	if choice == 1 {
		var s int
		fmt.Print("enter a side :")
		fmt.Scan(&s)
		area := countcube(s)
		fmt.Printf("the area of the cube is : %d\n", area)
	} else if choice == 2 {
		var r float64
		fmt.Print("enter the radius: ")
		fmt.Scan(&r)
		area := countcircle(r)
		fmt.Printf("the area of the circle is : %.2f\n", area)

	} else if choice == 3 {
		var rd float64
		var hg float64
		fmt.Print("enter the radius : ")
		fmt.Scan(&rd)
		fmt.Print("enter the height : ")
		fmt.Scan(&hg)
		ls := countcone(hg, rd)
		fmt.Println("the slaintheight of the cone is :", ls)
		area := countcone((rd), (hg))
		fmt.Printf("the area of a cone is : %2f\n", area)

	} else {
		fmt.Println("what do you choose bro??????")
	}
}
