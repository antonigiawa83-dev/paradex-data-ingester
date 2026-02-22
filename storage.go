package main

import "fmt"

func SimpanData(t Trade) {
	fmt.Printf(" [BROKER] Pushing price %.2f to buffer...\n", t.Price)
}
