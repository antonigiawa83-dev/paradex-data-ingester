package main

import "fmt"

// SimpanData menerima struct Trade yang baru (hanya Price dan Time)
func SimpanData(trade Trade) {
	// Kita hapus TradeID karena sudah tidak ada di struct utama
	fmt.Printf(" [STORAGE] Data tersimpan: Harga %.2f pada waktu %d\n", trade.Price, trade.Time)
}
