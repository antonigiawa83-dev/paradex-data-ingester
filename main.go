package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Struktur data utama - disamakan agar tidak error lagi
type TradeData struct {
	Price float64 `json:"price"`
	Time  int64   `json:"time"`
}

var (
	history     []TradeData
	historyLock sync.Mutex
)

func main() {
	// Mesin pembuat data simulasi
	go func() {
		for {
			historyLock.Lock()
			// Membuat harga acak di sekitar 67000
			now := time.Now()
			newData := TradeData{
				Price: 67000.0 + float64(now.Second()), 
				Time:  now.Unix(),
			}
			history = append(history, newData)
			
			// Simpan 100 data saja agar ringan
			if len(history) > 100 {
				history = history[1:]
			}
			historyLock.Unlock()
			time.Sleep(2 * time.Second)
		}
	}()

	// Handler API dengan Izin CORS
	http.HandleFunc("/api/history", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Content-Type", "application/json")

		historyLock.Lock()
		json.NewEncoder(w).Encode(history)
		historyLock.Unlock()
	})

	fmt.Println("=== Paradex Engine Fixed ===")
	fmt.Println("[SERVER] Berjalan di http://localhost:8080/api/history")
	
	server := &http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
