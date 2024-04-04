package main


import (
	"fmt"
	"time"
)



func main() {

	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			return 
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	go func(a, b int) bool {
		fmt.Println("a = ", a, " b = ", b)
		return true

	}(10,20)


	for {
		time.Sleep(time.Second)
	}

}