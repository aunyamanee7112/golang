package main

import "fmt"

func main() {
	/**** 🎃 Map 🎃 ****/
	menu := map[string]int{
		"strawberry cake":   10,
		"blueberry cake":    13,
		"choco banana cake": 8,
	}
	fmt.Println(menu)
	fmt.Println(menu["choco banana cake"])

	//loops
	for key, value := range menu {
		fmt.Println(value, "==>", key)
	}
	// ints as key type
	phoneList := map[int]string{
		830350854: "aun",
		970455468: "mom",
	}
	fmt.Println(phoneList)
	fmt.Println(phoneList[830350854])
	// object can change value by key
	phoneList[970455468] = "nuist"
	fmt.Println(phoneList)
}
