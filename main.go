package main

import (
	"fmt"
	"strings"
)

// func name (props) (values for return) {}
func getInitial(word string) (string, string) {
	upper := strings.ToUpper(word)
	names := strings.Split(upper, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}
func main() {
	/**** 🎃 Multiple Return Values 🎃 ****/
	/**** ✨ Multiple Return Values have to declare in func ,
	variable on call func ****/
	firstword1, secword1 := getInitial("Toshi")
	firstword2, secword2 := getInitial("Toshi Yamazaki")
	fmt.Printf("firstword1 is %v and secword2 is %v ", firstword1, secword1)
	fmt.Printf("firstword2 is %v and secword2 is %v ", firstword2, secword2)

}
