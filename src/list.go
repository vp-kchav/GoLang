package main

import (
	"container/list"
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {

	person1 := Person{"John", 44}
	person2 := Person{"Julia", 22}

	list := list.New()
	list.PushBack(person1)
	list.PushBack(person2)

	// Iterate the list
	for e := list.Front(); e != nil; e = e.Next() {
		//itemPerson := Person(e.Value.(Person))
		//fmt.Println(itemPerson.name)
		fmt.Println(e.Value.(Person).age, e.Value.(Person).name)
	}

}
