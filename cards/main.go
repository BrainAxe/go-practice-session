package main

import "fmt"

func main() {
		// var card string = "Ace of Spades"
		// card := newCard()
		// fmt.Println(card)	
		cards := []string{"Ace of Hearts", newCard()} // Slice
		cards = append(cards, "Six of Spades")

		for i, card := range cards {
			fmt.Println(i, card)
		}

}

func newCard() string {
		return "Five of Diamonds"
}