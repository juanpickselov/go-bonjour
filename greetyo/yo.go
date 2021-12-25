package main

import (
	"fmt"
	"log"
	"juanpickselov/greeter"
)

func main() {
	log.SetPrefix("greeter: ")
	log.SetFlags(0)

	message, err := greeter.Greet("Juan")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
