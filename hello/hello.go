package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	//message, error := greetings.Hello("Gladys")
	message, err := greetings.Hello("")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(message)
	}
}
