package main

import (
	"fmt"
)


func fibonacci(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			temp := x
			x = y 
			y = temp + y 
		case <- quit:
			return 
		}
	}

}
   

func main() {

	c := make(chan int)
	quit := make(chan int)

	defer close(c)
	defer close(quit)

	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<- c)
		}		
		quit <- 0
	}()

	fibonacci(c, quit)

}