package main

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()

	//how to create a slice
	cards := deck{"Ace of Diamonds", newCard()}

	//append returns new slice
	cards = append(cards, "Six of Spades")

	//for loop
	// for i, card := range cards {
	// 	fmt.Println(i, card)

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
