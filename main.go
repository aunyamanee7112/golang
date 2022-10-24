package main

import "fmt"

func main() {
	/**** 🎃 Booleans & Conditionals 🎃 ****/
	age := 45
	if age <= 45 {
		fmt.Println("this value less than 45")
	}
	// stagement inside the loops
	names := []string{"the8", "jun", "hoshi", "jeonghun", "wooji"}
	for index, value := range names {
		if index == 2 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v \n", index, value)
	}
}
