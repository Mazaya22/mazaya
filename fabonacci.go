package main

import "fmt"

func main() {

	Limit := 21
	a, b := 0, 1

	fmt.Print("Fibonacci : ")
	for a <= Limit {
		fmt.Print(a, ",")
		a, b = b, a+b
	}
}
