package main

import (
	"fmt"
	"time"
)



func main() {
	
	c := make(chan int, 3)
	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("go routine exit")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("go routine is running, sending element = ", i, "len(c) = ", len(c), "cap(c) = ", cap(c))
		}
	}()


	time.Sleep(2 * time.Second)


	for i := 0; i < 4; i++ {
		num := <- c
		fmt.Println("main routine is fetching num = ", num)
	}

	fmt.Println("main exit")


}