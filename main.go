package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Siapkan gudang data
	InitStorage()

	// 2. Jalankan mesin pencari data di latar belakang
	go StartConsumer()

	// 3. Buka pintu API dengan izin CORS agar grafik tidak kosong
	http.HandleFunc("/api/history", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		
		// Ambil data dari storage.go
		data := GetHistoryJSON()
		fmt.Fprintf(w, "%s", data)
	})

	fmt.Println("=== Paradex Ecosystem Connected ===")
	fmt.Println("[SERVER] API Aktif di http://localhost:8080/api/history")
	
	// Jalankan server
	http.ListenAndServe(":8080", nil)
}
