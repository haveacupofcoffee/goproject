package main

import (
	"fmt"
)


type SuperHumanInf interface{
	walk()
	fly()
	eat()
	getName() string
	getAge() int
}

type SuperHuman struct {
	name string
	age int
}

type SuperMan struct{
	SuperHuman
}


func (sman *SuperMan) walk() {
	fmt.Println("Superman walk")
}

func (sman *SuperMan) fly() {
	fmt.Println("Superman fly")
}

func (sman *SuperMan) eat() {
	fmt.Println("Superman eat")
}

func (sman *SuperMan) getName() string {
	return sman.name
}

func (sman *SuperMan) getAge() int {
	return sman.age
}

type SuperWoman struct {
	SuperHuman
}

func (swoman *SuperWoman) walk() {
	fmt.Println("SuperWoman walk")
}

func (swoman *SuperWoman) fly() {
 	fmt.Println("SuperWoman fly")
}

func (swoman *SuperWoman) eat() {
	fmt.Println("SuperWoman eat")
}

func (swoman *SuperWoman) getName() string {
	return swoman.name
}

func (swoman *SuperWoman) getAge() int {
	return swoman.age
}

func show(shuman SuperHumanInf) {
	fmt.Println("name = ", shuman.getName())
	fmt.Println("age = ", shuman.getAge())
}


func main() {
	var sman SuperMan
	sman.name = "Jason"
	sman.age = 33

	var swoman SuperWoman
	swoman.name = "Amelie"
	swoman.age = 27

	sman.fly()
	sman.eat()
	swoman.eat()

	show(&sman)
	show(&swoman)
}