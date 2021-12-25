package main

import (
	"fmt"
	"juanpickselov/greeter"
)

func main() {
	message := greeter.Greet("Juan")
	fmt.Println(message)
}

