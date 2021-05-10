package main

import ("fmt"
		"rsc.io/quote"
		"github.com/juanpickselov/go-bonjour/morestrings"
)

func main() {
	gsm := morestrings.ReverseRunes("accorD hondA")
	fmt.Println(gsm)
	fmt.Println(morestrings.ReverseRunes("corollA"))
	fmt.Println("Bonjour tout le monde!")
	fmt.Println(quote.Go())
}
