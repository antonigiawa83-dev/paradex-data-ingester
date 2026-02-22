package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http" // Package untuk membuat Server Web
	"time"
)

type Trade struct {
	TradeID   int     `json:"trade_id"`
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	Amount    float64 `json:"amount"`
	Side      string  `json:"side"`
	Timestamp string  `json:"timestamp"`
}

// Global variable untuk menyimpan data terakhir (simulasi database)
var lastTrade Trade

func main() {
	fmt.Println("=== Paradex API & Ingester System Starting ===")

	// 1. Jalankan Ingester di background agar data terus terupdate
	go func() {
		for i := 1; ; i++ {
			lastTrade = Trade{
				TradeID:   i,
				Symbol:    "BTC-USD",
				Price:     67000 + (rand.Float64() * 500),
				Amount:    rand.Float64() * 0.5,
				Side:      "buy",
				Timestamp: time.Now().Format(time.RFC3339),
			}
			time.Sleep(2 * time.Second)
		}
	}()

	// 2. Endpoint API: Browser akan mengambil data dari sini
	http.HandleFunc("/api/trade", getTradeHandler)

	fmt.Println(" [SERVER] API is running on http://localhost:8080/api/trade")
	
	// Jalankan Server di Port 8080
	http.ListenAndServe(":8080", nil)
}

func getTradeHandler(w http.ResponseWriter, r *http.Request) {
	// Memberitahu browser bahwa kita mengirim data JSON
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Agar bisa diakses dari Web mana saja
	
	json.NewEncoder(w).Encode(lastTrade)
}
