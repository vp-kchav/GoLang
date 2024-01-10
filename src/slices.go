package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	//To create an empty slice with non-zero length, use the builtin make.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println("get", s[5])

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	//sl1: [c d e]
	l := s[2:5]
	fmt.Println("sl1:", l)

	//sl2: [a b c d e]
	l = s[:5]
	fmt.Println("sl2:", l)

	//sl3: [c d e f]
	l = s[2:]
	fmt.Println("sl3:", l)

	//dcl: [g h i]
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// t == t2
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	//2d:  [[0] [1 2] [2 3 4]]
	fmt.Println("2d: ", twoD)
}
