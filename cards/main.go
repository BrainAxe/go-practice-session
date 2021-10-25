package main

func main() {
    // var card string = "Ace of Spades"
    // card := "Ace of Spades"
    cards := newDeck()
    cards.saveToFile("my_cards")
    
    // hand, remainingCards := deal(cards, 5)

    // hand.print()
    // remainingCards.print()
}