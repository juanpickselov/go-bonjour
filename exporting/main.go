package main

import (
	"fmt"
	"github.com/moficodes/go-crash-course/exporting/print"
)

func Export() int {
	return 42
}

func main() {
	fmt.Println(print.Number)
	fmt.Println(print.Data)
	fmt.Println("Data and Number printed because of the capital letters!")
	fmt.Println(Export())
}
