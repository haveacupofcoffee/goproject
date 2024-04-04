package main 

import (
	"fmt"
	"time"
)



func newTask() {

    i := 0
	for {
		fmt.Println("newTask i = ", i)
		i++
		time.Sleep(time.Second)
	}

}


func main() {

	go newTask()

	i := 0
	for {
		fmt.Println("main i = ", i)
		i++
		time.Sleep(time.Second)
	}

}