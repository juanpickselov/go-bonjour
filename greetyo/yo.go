package main

import (
	"fmt"
	"juanpickselov/greeter"
	"log"
)

func main() {
	log.SetPrefix("greeter: ")
	log.SetFlags(0)

	names := []string{"Juan", "Sofia", "Anastasia"}

	message, err := greeter.Greets(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
