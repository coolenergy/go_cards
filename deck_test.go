package main

import "testing"

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 16 {
		panic("Wrong number of cards in the deck")
	}
}
