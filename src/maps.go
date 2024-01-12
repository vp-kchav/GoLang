package main

import (
	"fmt"
	"maps"
)

func main() {
	myMap := make(map[string]int)
	myMap["k1"] = 12
	myMap["k2"] = 18
	fmt.Println("printMap", myMap)

	v1 := myMap["k1"]
	fmt.Println("v1 Value: ", v1)

	v3 := myMap["k3"]
	fmt.Println("v3:", v3)

	// result: 2
	fmt.Println("len:", len(myMap))

	//delete keyword is to delete any key from map
	delete(myMap, "k2")
	fmt.Println("map:", myMap)

	//result 0 because it already deleted
	v2 := myMap["k2"]
	fmt.Println("v2:", v2)

	fmt.Println("len:", len(myMap))

	//clear keyword is to clear the map
	clear(myMap)
	fmt.Println("map:", myMap)
	fmt.Println("len:", len(myMap))

	myMap["k3"] = 12
	// '_,' is to check if the value is presented with the specific key provided.
	_, prs := myMap["k3"]
	_, prs1 := myMap["k2"]
	//return true
	fmt.Println("prs:", prs)
	//return false
	fmt.Println("prs1:", prs1)

	// how to initialize map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}

	//using maps pacakge for many other useful libs
	n3 := map[string]int{"foo": 1, "bar": 2, "boo": 3}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
	if maps.Equal(n, n3) {
		fmt.Println("n == n2")
	} else {
		fmt.Println("n != n3")
	}
}
