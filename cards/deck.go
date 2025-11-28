package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//*Create a new type of deck
//*which is a slice of strings

// *kind of an alias
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	//*converting it into type []string
	return strings.Join([]string(d), ",")
	//*basically what we did here was to join the string using the , as seperator
	//* join(a [] string)
	//*joining the array

}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	//*if everything is correct the value of err == nil
	if err != nil {
		//apply common sense
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return (s) //Ace of spades,Two of Spades
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {

		newPosition := r.Intn(len(d) - 1)
		//*swap swap
		d[i], d[newPosition] = d[newPosition], d[i]
		//gen random math say t
		//or we can be clever with it

	}
}
