package main

import "fmt"

//*now we start with the print function,
//*this is where we would take it as an argument

func print(d data) {
	//*what do we do here
	for i, k := range d {

		if k == 1 {
			fmt.Printf("index %d is odd \n", i)
		} else {
			fmt.Printf("index %d is even\n", i)
		}

	}
}
