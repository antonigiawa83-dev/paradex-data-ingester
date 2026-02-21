package main

import (
	"fmt"
	"time"
)

// RunConsumer mensimulasikan proses pengambilan data dari Broker ke Database
func RunConsumer() {
	fmt.Println(" [CONSUMER] Worker started: Waiting for trades from Broker...")

	// Di dunia nyata, ini akan menggunakan library Kafka consumer
	for {
		fmt.Println(" [CONSUMER] Consuming 1 batch of trades from Topic 'paradex-trades'...")
		
		// Simulasi proses "Batch Insert" ke Clickhouse (lebih efisien daripada satu per satu)
		time.Sleep(3 * time.Second) 
		
		fmt.Println(" [DB-SYNC] Batch successfully written to Clickhouse. Index updated.")
	}
}
