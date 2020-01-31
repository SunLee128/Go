package main

import "fmt"

type deck []string

//receiver - any variable of type 'deck' gets access to the 'print' method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
