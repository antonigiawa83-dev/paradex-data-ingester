package main

import (
	"encoding/json"
	"sync"
)

type Trade struct {
	Price float64 `json:"price"`
	Time  int64   `json:"time"`
}

var (
	historyData []Trade
	mu          sync.Mutex
)

func InitStorage() {
	historyData = make([]Trade, 0)
}

func SaveTrade(price float64, timestamp int64) {
	mu.Lock()
	defer mu.Unlock()
	historyData = append(historyData, Trade{Price: price, Time: timestamp})
	if len(historyData) > 100 {
		historyData = historyData[1:]
	}
}

func GetHistory() string {
	mu.Lock()
	defer mu.Unlock()
	bytes, _ := json.Marshal(historyData)
	return string(bytes)
}
