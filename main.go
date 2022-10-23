package main

import "fmt"

func main() {
	/**** 🎃 Loops 🎃 ****/
	x := 5
	// ❌ not recommend
	for x < 5 {
		// fmt.Print("value of x ", x)
		// x++
	}
	// ✅ recommend
	for i := 0; i < 10; i++ {
		// fmt.Println(i)
	}

	// example
	list := []string{"aun", "nui", "noey"}
	for i := 0; i < len(list); i++ {
		// fmt.Println(list[i])
	}

	for index, value := range list {
		fmt.Printf("this index is %v and this value is %v \n", index, value)
	}
	// need only value
	for _, value := range list {
		fmt.Printf("value is %v \n", value)
	}

}
