package main

import (
	"fmt"
	"strconv"
)

func main() {

	var str string = "hello"
	var n int64 = 11

	n, _ = strconv.ParseInt(str, 10, 64)

	fmt.Printf("n type is %T, n = %d\n", n, n)

}