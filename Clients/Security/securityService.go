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
		go emitSecurityLog(i)
	}
	select {}
}

func emitSecurityLog(id int) {
	for {
		log := map[string]interface{}{
			"timestamp": time.Now().UTC(),
			"level":     randomLevel(),
			"source":    "security-service",
			"message":   "Firewall rule triggered",
			"user_id":   rand.Intn(500),
		}
		send(log)
		time.Sleep(time.Second * time.Duration(rand.Intn(4)+2))
	}
}

func send(log map[string]interface{}) {
	body, _ := json.Marshal(log)
	http.Post("http://localhost:8001/ingest", "application/json", bytes.NewBuffer(body))
}

func randomLevel() string {
	return []string{"info", "warn", "error"}[rand.Intn(3)]
}
