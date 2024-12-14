package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("punch to start!!!!")
	countdown := 10

	fmt.Scan(&countdown)
	fmt.Println("time is over in")

	for c := countdown; c >= 1; c-- {
		fmt.Printf("%d \n", c)

		time.Sleep(1 * time.Second)
	}

	fmt.Println("diririring!!!!!")
}
