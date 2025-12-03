package main

import (
	"fmt"
	"time"
)

func bug() {
	links := []string{"a", "b", "c"}
	for _, l := range links {
		fmt.Println("value of l at the time of creation of go-routine", l)
		go func() {
			fmt.Println("value captured at time of execution", l) // captures l (the loop variable)
		}()
	}
	time.Sleep(100 * time.Millisecond)

}
