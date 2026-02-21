package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
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

func main() {
	fmt.Println("=== Paradex Data Pipeline System Starting ===")

	// 1. Jalankan Consumer sebagai background process (Worker)
	// Ini akan mengambil data dari 'buffer' secara asinkron
	go RunConsumer()

	fmt.Println(" [MAIN] Ingester Engine is active...")

	for i := 1; ; i++ {
		side := "buy"
		if rand.Intn(2) == 0 {
			side = "sell"
		}

		trade := Trade{
			TradeID:   i,
			Symbol:    "BTC-USD",
			Price:     95000 + (rand.Float64() * 500),
			Amount:    rand.Float64() * 0.5,
			Side:      side,
			Timestamp: time.Now().Format(time.RFC3339),
		}

		// LOGIKA MONITORING
		if trade.Price > 95350 {
			fmt.Printf("⚠️  ALERT: High Volatility! Price: %.2f\n", trade.Price)
		}

		// Tampilkan log data mentah
		jsonData, _ := json.Marshal(trade)
		fmt.Println(" [INGEST] Received: ", string(jsonData))

		// 2. Kirim data ke Broker (Storage Layer)
		SimpanData(trade)

		time.Sleep(2 * time.Second) 
	}
}
