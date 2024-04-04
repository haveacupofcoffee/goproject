package main

import (
	"fmt"
)


func myFunc(arg interface{}) {
	fmt.Println("myFunc is called ....")
	fmt.Println("arg = ", arg)

	value, ok := arg.(string)
	fmt.Printf("vlaue = %v, ok = %v\n", value, ok)
}

type Book struct {
	name string
}

func main() {

	book := Book{"go lang"}
	book1 := &book
	fmt.Println(book)
	fmt.Println(book1.name)
	myFunc(book)
	myFunc(21)
	myFunc("abc")
	myFunc(3.14)

}