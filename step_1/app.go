package main

import (
	"fmt"

	"org.gfoo/go-demo-1/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello world from BestMAscot " + mascot.BestMascot())
	fmt.Println(quote.Go())
}
