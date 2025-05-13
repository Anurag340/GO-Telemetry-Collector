package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go emitCloudLog(i)
	}
	select {}
}

func emitCloudLog(id int) {
	for {
		log := map[string]interface{}{
			"timestamp": time.Now().UTC(),
			"level":     randomLevel(),
			"source":    "cloud-service",
			"message":   "VM instance scaled",
			"user_id":   rand.Intn(800),
			"amount":    float64(rand.Intn(5000)) / 100,
		}
		send(log)
		time.Sleep(time.Second * time.Duration(rand.Intn(5)+1))
	}
}

func send(log map[string]interface{}) {
	body, _ := json.Marshal(log)
	http.Post("http://localhost:8001/ingest", "application/json", bytes.NewBuffer(body))
}

func randomLevel() string {
	return []string{"info", "warn", "error"}[rand.Intn(3)]
}
