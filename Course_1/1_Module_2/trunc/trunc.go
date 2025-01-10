package main

import (
	"fmt"
	
)

func main() {
	var input float64

	fmt.Print("Enter flaoting number: ")
	_,err := fmt.Scan(&input)

	if(err != nil){
		return
	}

	truncated := int(input)
	fmt.Printf("The truncated integer is: %d\n", truncated)
}

