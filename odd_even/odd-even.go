package main

//*making a function for odd and even

func (d *data) oddEven() {
	//*we will just populate the array with odd even
	//*we swap the values accordingly

	for i, n := range *d {
		if n%2 == 1 {
			(*d)[i] = 1
		} else {
			(*d)[i] = 0
		}
	}
}
