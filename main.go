package main

func main() {
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Ace of Diamonds")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
