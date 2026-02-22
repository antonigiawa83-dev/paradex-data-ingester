package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Struct Trade ini harus SAMA di semua file agar tidak error
type Trade struct {
	Price float64 `json:"price"`
	Time  int64   `json:"time"`
}

// Variabel untuk menyimpan riwayat data grafik
var tradeHistory []Trade

func main() {
	fmt.Println("=== Paradex Data Pipeline System Starting ===")

	// 1. Jalankan Ingester & Storage secara paralel
	go func() {
		for {
			// Membuat data trade baru (Hanya Price dan Time)
			newTrade := Trade{
				Price: 67000 + (rand.Float64() * 500),
				Time:  time.Now().Unix(),
			}

			// Simpan ke riwayat untuk ditampilkan di grafik (maks 50 data)
			tradeHistory = append(tradeHistory, newTrade)
			if len(tradeHistory) > 50 {
				tradeHistory = tradeHistory[1:]
			}

			// MEMANGGIL FUNGSI DARI storage.go
			// Pastikan storage.go kamu sudah menggunakan: func SimpanData(t Trade)
			SimpanData(newTrade)

			time.Sleep(2 * time.Second)
		}
	}()

	// 2. MEMANGGIL FUNGSI DARI consumer.go
	go RunConsumer()

	// 3. API Endpoint untuk index.html (Grafik)
	http.HandleFunc("/api/history", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(tradeHistory)
	})

	fmt.Println(" [SERVER] API aktif di http://localhost:8080/api/history")
	
	// Jalankan server web
	http.ListenAndServe(":8080", nil)
}
