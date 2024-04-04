package main

import (
	"reflect"
	"fmt"
)

func reflectArgs(arg interface{}) {
	fmt.Println("Type is : " , reflect.TypeOf(arg))
	fmt.Println("Value is: " , reflect.ValueOf(arg))
}




func main() {
	var num float64 = 1.234
	reflectArgs(num)

}