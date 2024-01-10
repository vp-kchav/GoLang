package main

import (
	"fmt"
)

func main() {

	if 72%2 == 0 {
		fmt.Println("I am even")
	} else {
		fmt.Println("I am odd")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num := 11; num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
