package main

type data []int

func (d *data) feed() {
	for i := 1; i <= 10; i++ {
		*d = append(*d, i)
	}
}
