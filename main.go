package main

import "fmt"

func main() {
	//Arrays
	number := [3]int{4, 50, 60}
	name := [2]string{"hi", "anyoung"}
	fmt.Println(number, len(number), "number & lenght")
	fmt.Println(name, len(name), "name  & lenght")
	// slice use array under the hood
	score := []int{100, 50, 70}
	// 🎃 not working
	score[2] = 35
	// 🧚🏻 it's work
	score = append(score, 60)
	fmt.Println(score, len(score))
	// slice range
	fruities := []string{"apple", "pineapple", "dragonfruit", "giwi"}
	// 🎃 start and get me up to position end
	// but not including end position to arr
	rangefruities1 := fruities[1:3]
	// 🎃 start - the end of arr
	rangefruities2 := fruities[2:]
	// 🎃 the start of arr - before end position
	rangefruities3 := fruities[:3]
	fmt.Println(rangefruities1, "range fruit1")
	fmt.Println(rangefruities2, "range fruit2")
	fmt.Println(rangefruities3, "range fruit3")

}
