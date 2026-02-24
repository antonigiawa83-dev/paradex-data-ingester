package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Inisialisasi Penyimpanan (dari storage.go)
	InitStorage()

	// 2. Jalankan Pengambil Data (dari consumer.go)
	go StartConsumer()

	// 3. Handler API untuk Grafik dengan izin CORS
	http.HandleFunc("/api/history", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		
		data := GetHistory() // Ambil data dari storage.go
		fmt.Fprintf(w, "%s", data)
	})

	fmt.Println("=== Paradex System Connected & Harmonized ===")
	fmt.Println("[SERVER] API siap di http://localhost:8080/api/history")
	http.ListenAndServe(":8080", nil)
}
