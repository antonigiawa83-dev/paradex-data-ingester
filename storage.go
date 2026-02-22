package main

import "fmt"

// Pastikan parameter 't' menggunakan tipe Trade yang baru (Price & Time saja)
func SimpanData(t Trade) {
	fmt.Printf(" [BROKER] Recording price: %.2f at timestamp: %d\n", t.Price, t.Time)
}
