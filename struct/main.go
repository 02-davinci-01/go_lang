package main

import "fmt"

type fruit interface {
	getFruit() int
}

type orange string
type apple string

func main() {
	newOrange := orange("")
	newApple := apple("")
	printingFruit(newOrange)
	printingFruit(newApple)
}

func (o orange) getFruit() int {
	return 1
}
func (a apple) getFruit() int {
	return 2
}

func printingFruit(f fruit) {
	fmt.Printf("this is the number of fruit %v", f.getFruit())
}
