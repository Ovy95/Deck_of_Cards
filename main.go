package main

func main() {
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Five of Hearts")
	cards.print()

}

func newCard() string {
	return "Nine of Clubs"
}
