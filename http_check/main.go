package main

import (
	"fmt"
	"os"
)

func osWrite() {

	//gives me the full output

	//opening the file
	fptr, error := os.Open(os.Args[1])
	if error != nil {
		fmt.Println("file error")
		os.Exit(1)
	}
	fmt.Println(fptr) //we have a pointer of type File
	//let's see if we can write that or not
	//this fptr has a read function associtaed with it
	//can we use the writer function?
	conWrite := make([]byte, 20)
	byt, error := fptr.Read(conWrite)
	fmt.Println(byt)

	if error != nil {
		fmt.Println("reading error")
	}

	type str []string
	fmt.Println(string(conWrite))

}
