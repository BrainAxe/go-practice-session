package main

func main() {
    // var card string = "Ace of Spades"
    // card := "Ace of Spades"
    cards := newDeck()
    cards.shuffle()
    cards.print()
    // cards.saveToFile("my_cards")

    // cards := newDeckFromFile("my_cards")
    // cards.print()
    
    // hand, remainingCards := deal(cards, 5)

    // hand.print()
    // remainingCards.print()
}