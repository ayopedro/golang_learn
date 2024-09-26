package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	// // cards.saveToFile("my_cards.txt")

	// // cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()

	// practice with slices
	// numbers := []int{}

	// for i := 0; i <= 10; i++ {
	// 	numbers = append(numbers, i)
	// }

	// for _, number := range numbers {
	// 	if number%2 == 0 {
	// 		fmt.Println(number, " is even")
	// 	} else {
	// 		fmt.Println(number, " is odd")

	// 	}
	// }
}
