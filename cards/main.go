package main

import "fmt"

// import "fmt"

func main() {
	cards := newDeck()
	cards.print()

	//*seperation
	fmt.Println("******************")
	cards.shuffle()

	cards.print()

}
