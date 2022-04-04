package main

import "fmt"

//Declare Games Struct
type game struct {
	title string
	price float64
}

//Create print method
func (game game) printPrice() {
	fmt.Printf("Title: %v  Price: $%.2f\n", game.title, game.price)
}