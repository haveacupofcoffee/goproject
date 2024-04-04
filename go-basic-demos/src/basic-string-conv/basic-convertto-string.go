package main

import (
	"fmt"
	"strconv"
)


func main() {
    ConvertMethod1()
	ConvertMethod2()
}

func ConvertMethod1() {
	fmt.Println("Convert Method1 with fmt.Sprintf()")
	var n int32 = 1280
	var f float32 = 2.3458
	var b bool = true
	var c byte = 'a'

	var str string

	str = fmt.Sprintf("%d", n)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	str = fmt.Sprintf("%f", f)
	fmt.Printf("str type is %T, str = %q\n", str, str)
	
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	str = fmt.Sprintf("%c", c)
	fmt.Printf("str type is %T, str = %q\n", str, str)
}

func ConvertMethod2() {
	fmt.Println("Covert Method2 with strconv.FormatXXX()")

	var n int32 = 1024
	var f float32 = 2.3456
	var b bool = true

	var str string

	str = strconv.FormatInt(int64(n), 10)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	str = strconv.FormatFloat(float64(f), 'f', 5, 64)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	str = strconv.FormatBool(b)
	fmt.Printf("str type is %T, str = %q\n", str, str)


}