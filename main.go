package main

func main() {
	cards := newDeck()
	// cards = append(cards, "Six of Spades")
	//save to file
	cards.shuffle()
	cards.print()

	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()
}
