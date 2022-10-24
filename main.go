package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("hi good morning %v \n", n)
}
func sayGoodbye(n string) {
	fmt.Printf("hi goodbye %v \n", n)
}
func renderName(n []string, f func(string)) {
	for _, v := range n {
		f(v)
		sayGoodbye(v)
	}
}
func calCircle(r float64) float64 {
	return math.Pi * r * r
	// assign type on function again because value return is over float64
}

func main() {
	/**** 🎃 Using Functions 🎃 ****/

	// normal test
	sayGreeting("aun")
	sayGoodbye("aun")
	// test with func loop
	friends := []string{"aun", "elle", "noey", "nui", "ky", "mhon"}
	renderName(friends, sayGreeting)
	// finding area circle
	circle := calCircle(10.7)
	fmt.Printf("this circle area is %0.3f", circle)
}
