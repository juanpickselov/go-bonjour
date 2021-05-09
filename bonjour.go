package main

import ("fmt"
		"rsc.io/quote"
		"github.com/juanpickselov/go-bonjour/morestrings"
)

func main() {
	gsm := morestrings.ReverseRunes("accorD hondA")

	fmt.Println("\n" + quote.Go(),"\nBonjour tout le monde!", "\n" + gsm)
}
