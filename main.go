package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	/**** 🎃 The standard Library 🎃 ****/

	//strings
	greeting := "hello ginger breads"
	// 🌝 check text in variable
	fmt.Println(strings.Contains(greeting, "aun"))           // get false ❌
	fmt.Println(strings.Contains(greeting, "hello"))         // get true  ✅
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) // replace but not change value in via
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "g"))
	fmt.Println(strings.Split(greeting, " "))
	fmt.Println(strings.Split(greeting, "ginger"))

	//sort
	names := []string{"may", "nui", "noey", "aun", "elle"}
	sort.Strings(names)
	search := sort.SearchStrings(names, "may")
	fmt.Println(names, search)
}
