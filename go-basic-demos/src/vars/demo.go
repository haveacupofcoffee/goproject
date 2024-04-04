package main

import "fmt"

var (
	n1, n2, n3 = 10, 20, "hello"
)
func main() {
	var num int = 10
	fmt.Println("num=", num)

	fmt.Println("n1=", n1, "n2=",n2, "n3=",n3)

	var n4 int32 = 12
	var n5 int8 = int8(n4) + 127
	//var n6 int8 = int8(n4) + 128

	fmt.Println("n4=", n4)


}