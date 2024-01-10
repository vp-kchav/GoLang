package main

import (
	"fmt"
)

func main() {
	i := 1
	for i < 3 {
		fmt.Print(i, "\n")
		i++
	}

	for j := 1; j < 5; j++ {
		fmt.Print(j, "\n")
	}

	for j := 1; j < 10; j++ {
		if j%2 == 0 {
			fmt.Print(j, "\n")
		}
	}
}
