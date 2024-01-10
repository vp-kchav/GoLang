package main

import (
	"fmt"
)

type Book struct {
	name   string
	author string
	page   int
}

func (book Book) printDetails() {
	fmt.Printf("The book '%s' was written by '%s' has '%d' pages. \n", book.name, book.author, book.page)
}

func printDetails(book Book) {
	fmt.Printf("The book '%s' was written by '%s' has '%d' pages. \n", book.name, book.author, book.page)
}

func printMe(name, author string, page int) {
	fmt.Printf("The book '%s' was written by '%s' has '%d' pages. \n", name, author, page)
}

func plus(x, y int) int {
	return (x + y)
}

func main() {
	book1 := Book{"hello", "Kimtey", 122}
	book1.printDetails()

	book2 := Book{"hello2", "Alvin", 324}
	book2.printDetails()

	book3 := Book{}
	book3.author = "Sethika"
	book3.name = "baby"
	book3.page = 543
	book3.printDetails()

	fmt.Println("next level of method")

	printDetails(book1)
	printDetails(book2)
	printDetails(book3)

	printMe("Xing", "crying", 12)

	fmt.Println("return function")
	var sumPlus = plus(7, 8)
	fmt.Println(sumPlus)
}
