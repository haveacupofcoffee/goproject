package main

import (
	"fmt"
)


type Hero struct {
	name string
	age int
	level int
}

func (this Hero) GetName() string {
	return this.name
}

func (this *Hero) SetName(name string) {
	this.name = name
}

func (this Hero) show() {
	fmt.Println("name = ", this.GetName())
	fmt.Println("age = " , this.age)
	fmt.Println("level = " , this.level)

}


// a copy of Hero
// func (this Hero) GetName() string {
// 	return this.name
// }

// func (this Hero) SetName(name string) {
// 	this.name = name
// }

// func (this Hero) show() {
// 	fmt.Println("name = " , this.name)
// 	fmt.Println("adge = " , this.age)
// 	fmt.Println("level = ", this.level)
// }


func main() {
	hero := Hero{name : "hello", age : 10, level : 1}
    
	hero.show()

	hero.SetName("world")

	hero.show()

}