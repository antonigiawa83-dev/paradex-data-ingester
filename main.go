package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Trade struct {
	Price     float64 `json:"price"`
	Time      int64   `json:"time"`
}

// Menyimpan riwayat data untuk grafik
var tradeHistory []Trade

func main() {
	fmt.Println("=== Paradex Chart Engine Starting ===")

	go func() {
		for {
			newTrade := Trade{
				Price: 67000 + (rand.Float64() * 500),
				Time:  time.Now().Unix(),
			}
			// Simpan ke riwayat
			tradeHistory = append(tradeHistory, newTrade)
			
			// Batasi hanya 50 data terakhir agar tidak berat
			if len(tradeHistory) > 50 {
				tradeHistory = tradeHistory[1:]
			}
			time.Sleep(2 * time.Second)
		}
	}()

	http.HandleFunc("/api/history", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(tradeHistory)
	})

	fmt.Println(" [SERVER] Running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
