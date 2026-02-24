package main

import (
	"encoding/json"
	"sync"
)

// Definisi Trade harus ada di sini agar dikenal oleh consumer.go dan main.go
type Trade struct {
	Price float64 `json:"price"`
	Time  int64   `json:"time"`
}

var (
	historyData []Trade
	mu          sync.Mutex
)

// Fungsi untuk menyiapkan tempat penyimpanan
func InitStorage() {
	mu.Lock()
	defer mu.Unlock()
	historyData = make([]Trade, 0)
}

// Fungsi untuk menyimpan data (dipanggil oleh consumer.go)
func SaveTrade(price float64, timestamp int64) {
	mu.Lock()
	defer mu.Unlock()
	
	newData := Trade{Price: price, Time: timestamp}
	historyData = append(historyData, newData)
	
	// Batasi hanya 100 data agar tidak lambat
	if len(historyData) > 100 {
		historyData = historyData[1:]
	}
}

// Fungsi untuk mengambil data dalam format JSON (dipanggil oleh main.go)
func GetHistoryJSON() string {
	mu.Lock()
	defer mu.Unlock()
	
	bytes, err := json.Marshal(historyData)
	if err != nil {
		return "[]"
	}
	return string(bytes)
}
