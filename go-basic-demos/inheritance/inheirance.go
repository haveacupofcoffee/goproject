package main

import (
	"fmt"
)

type Human struct {
	name string 
	age int
}

func (this *Human) Eat() {
	fmt.Println("Human eat.....")
}

func (this *Human) Walk() {
	fmt.Println("Human walk......")
}

type SuperHuman struct {
	Human
	level int
}

func (this *SuperHuman) Fly() {
	fmt.Println("SuperHuman fly....")
}

func (this *SuperHuman) Eat() {
	fmt.Println("SuperHuman eat.....")
}

func (this *SuperHuman) print() {
	fmt.Println("name = " , this.name)
	fmt.Println("age = ", this.age)
	fmt.Println("level = ", this.level)
	
}

func main() {
	
	h := Human{"Jason", 20}

	h.Walk()
	h.Eat()

	//s := SuperHuman{Human{"Mark", 21}, 20}
	var s SuperHuman
	s.name = "Mark"
	s.age = 21
	s.level = 20

	s.Fly()
	s.Walk()
	s.Eat()
	s.print()

}