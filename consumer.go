package main

import (
	"time"
)

func StartConsumer() {
	for {
		// Simulasi pengambilan data harga
		price := 67000.0 + float64(time.Now().Second())
		timestamp := time.Now().Unix()
		
		// Simpan ke storage.go
		SaveTrade(price, timestamp)
		
		time.Sleep(2 * time.Second)
	}
}
