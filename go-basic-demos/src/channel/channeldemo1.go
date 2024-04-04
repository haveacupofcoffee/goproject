package main

import (
	"fmt"
)


func main () {

	c := make(chan int)

	go func() {
		defer fmt.Println("go routine exit")
		
		c <- 666
		fmt.Println("go routine")
	}()



	num := <- c
    fmt.Println("num = ", num)
	fmt.Println("main exit")


}