package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// Trade mewakili data transaksi di Paradex
type Trade struct {
	TradeID   int     `json:"trade_id"`
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	Amount    float64 `json:"amount"`
	Side      string  `json:"side"` // "buy" atau "sell"
	Timestamp string  `json:"timestamp"`
}

func main() {
	fmt.Println("Paradex Real-time Ingester Engine is running...")

	// Simulasi stream data tanpa henti
	for i := 1; ; i++ {
		// Logika acak untuk menentukan buy/sell
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

		// Mengubah struct menjadi JSON (Data Serialization)
		data, _ := json.Marshal(trade)
		fmt.Println(string(data))

		// Di industri nyata, data ini dikirim ke Kafka atau Clickhouse
		time.Sleep(2 * time.Second) 
	}
}
