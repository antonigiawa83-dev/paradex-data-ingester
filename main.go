package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Struktur data untuk grafik
type TradeData struct {
	Price float64 `json:"price"`
	Time  int64   `json:"time"`
}

var (
	history     []TradeData
	historyLock sync.Mutex
)

func main() {
	// Goroutine untuk simulasi data (pengganti consumer jika belum konek ke bursa)
	go func() {
		for {
			historyLock.Lock()
			newData := TradeData{
				Price: 67000.0 + (time.Now().Unix() % 100).ToFloat64(), // Harga simulasi sesuai screenshot
				Time:  time.Now().Unix(),
			}
			history = append(history, newData)
			// Batasi data agar tidak lemot (simpan 100 data terakhir)
			if len(history) > 100 {
				history = history[1:]
			}
			fmt.Printf("[STORAGE] Data tersimpan: Harga %.2f pada waktu %d\n", newData.Price, newData.Time)
			historyLock.Unlock()
			time.Sleep(2 * time.Second)
		}
	}()

	// Endpoint API yang dipanggil oleh index.html
	http.HandleFunc("/api/history", func(w http.ResponseWriter, r *http.Request) {
		// --- SURAT IZIN (CORS) DIMULAI ---
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Content-Type", "application/json")
		// --- SURAT IZIN (CORS) SELESAI ---

		historyLock.Lock()
		json.NewEncoder(w).Encode(history)
		historyLock.Unlock()
	})

	fmt.Println("=== Paradex Data Pipeline System Starting ===")
	fmt.Println("[SERVER] API aktif di http://localhost:8080/api/history")
	http.ListenAndServe(":8080", nil)
}

// Fungsi pembantu untuk konversi
func (i int64) ToFloat64() float64 {
	return float64(i)
}
