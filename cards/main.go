package main

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()

	//how to create a slice
	cards := newDeck()
	hand, remainingCard := deal(cards, 5)
	//multiple return
	hand.print()
	remainingCard.print()
}
