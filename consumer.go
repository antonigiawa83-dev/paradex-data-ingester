package main

import (
	"time"
)

func StartConsumer() {
	for {
		// Simulasi harga sekitar 67000 seperti di terminal kamu
		price := 67000.0 + float64(time.Now().Second())
		timestamp := time.Now().Unix()
		
		// Kirim data ke storage.go
		SaveTrade(price, timestamp)
		
		time.Sleep(2 * time.Second)
	}
}
