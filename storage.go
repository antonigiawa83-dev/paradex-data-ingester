package main

import (
	"encoding/json"
	"sync"
)

// Struktur data utama yang digunakan bersama
type Trade struct {
	Price float64 `json:"price"`
	Time  int64   `json:"time"`
}

var (
	historyData []Trade
	mu          sync.Mutex
)

// Inisialisasi awal
func InitStorage() {
	historyData = make([]Trade, 0)
}

// Menyimpan data baru (digunakan oleh consumer.go)
func SaveTrade(price float64, timestamp int64) {
	mu.Lock()
	defer mu.Unlock()
	historyData = append(historyData, Trade{Price: price, Time: timestamp})
	if len(historyData) > 100 {
		historyData = historyData[1:]
	}
}

// Mengambil data untuk API (digunakan oleh main.go)
func GetHistoryJSON() string {
	mu.Lock()
	defer mu.Unlock()
	bytes, _ := json.Marshal(historyData)
	return string(bytes)
}
