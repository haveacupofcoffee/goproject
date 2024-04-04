package main

import (
	"encoding/json"
	"fmt"
)


type Movie struct {
	Title string `json:"nexus3.title"`
	Year int `json:"nexus3.year"`
    TicketPrice int `json:"nexus3.price"`
	Actors []string `json:"nexus3.actors"`
}


func main() {
	movie := Movie{"Frineds", 1998, 10, []string{"Ross","Richel","Joey"}}

	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("convert failed.", err)
		return
	}

	fmt.Printf("jsonStr  = %s ", jsonStr)

	jsonStr = []byte(`{"title":"The Big Bang Theory","year":2005,"price":10,"actors":["Lenord","Harwod","Sheldon"]}`)

	movie2 := Movie{}
	
	err = json.Unmarshal(jsonStr, &movie2)

	if err != nil {
		fmt.Println("unmarshal failed.", err)
		return
	}

	fmt.Printf("moive = %v\n", movie2)
	
}