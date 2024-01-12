package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	//sum: 9
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			//index: 1
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		//a -> apple
		//b -> banana
		fmt.Printf("%s -> %s\n", k, v)
	}

	//key: a
	//key: b
	for k := range kvs {
		fmt.Println("key:", k)
	}

	//first paramenter is always key
	for v1, k1 := range kvs {
		fmt.Println("key use v:", k1)
		fmt.Println("key use v:", v1)
	}

	for i, c := range "goddffd" {
		fmt.Println("go", i, c)
	}

}
