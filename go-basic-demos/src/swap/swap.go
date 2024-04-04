package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20

	fmt.Println("before swap: a = ", a , "b = ", b)

	swap(&a, &b)

	fmt.Println("After swap: a = ", a, "b = ", b)
}

func swap ( a *int , b *int) {
	temp := *a 
	*a = *b 
	*b = temp
}

//wrong example 
// func swap(a int , b int) {
// 	temp := a
// 	a = b
// 	b = temp
// }