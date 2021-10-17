package main

import (
	"fmt"

	"github.com/juanpickselov/go-bonjour/morestrings"
	"rsc.io/quote"
)

func main() {
	gsm := morestrings.ReverseRunes("accorD hondA")
	fmt.Println(gsm)
	fmt.Println(morestrings.ReverseRunes("corollA"))
	fmt.Println("Bonjour VIM!")
	fmt.Println("你 好")
	fmt.Println(quote.Go())
}
