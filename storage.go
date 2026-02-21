package main

import (
	"fmt"
	"time"
)

// SimpanData sekarang mensimulasikan pengiriman ke Message Broker (Kafka/Redpanda)
func SimpanData(trade Trade) {
	fmt.Printf(" [BROKER] Mengirim transaksi ID %d ke Topic 'paradex-trades'...\n", trade.TradeID)
	
	// Simulasi sedikit delay jaringan ke broker
	time.Sleep(10 * time.Millisecond) 
	
	fmt.Printf(" [SUCCESS] Data aman di buffer. Database tetap stabil.\n")
}
