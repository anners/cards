package main

import "fmt"

func main() {
	//var card string = "Ace of Spaces"
	cards := newDeck()

	hand, remainigCards := deal(cards, 5)

	hand.print()
	fmt.Println("Remaining Cards")
	remainigCards.print()

	cards.saveToFile("mydeck")

}
