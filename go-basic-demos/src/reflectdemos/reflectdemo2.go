package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	ISBN string
	Name string
	Author string
}


func (b Book) Read() {
	fmt.Println("Book is read ....")
} 

func getAllFiledsAndMethods(arg interface{}) {
	inputType := reflect.TypeOf(arg)
	fmt.Println("type is: ", inputType.Name() )

	inputValue := reflect.ValueOf(arg)
	fmt.Println("value is:", inputValue)

	fmt.Println("Each field : ") 

	for i :=0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		fmt.Printf("%s : %v = %v\n", field.Name, field.Type, inputValue.Field(i).Interface())
	}

	fmt.Println("Mthod num : ", inputType.NumMethod())

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s : %v", m.Name, m.Type)
	}

}


func main() {
	book := Book{"1234", "One Punch Man", "unknown"}
	getAllFiledsAndMethods(book)
}