package main

import "fmt"

// SimpanData simulasi pengiriman data ke database OLAP
func SimpanData(trade Trade) {
	fmt.Printf(" [DB-SAVE] Berhasil menyimpan transaksi ID %d ke Clickhouse\n", trade.TradeID)
}
