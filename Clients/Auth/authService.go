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
		go emitAuthLog(i)
	}
	select {} // Keep alive
}

func emitAuthLog(id int) {
	for {
		log := map[string]interface{}{
			"timestamp": time.Now().UTC(),
			"level":     randomLevel(),
			"source":    "auth-service",
			"message":   "User login attempt",
			"user_id":   rand.Intn(1000),
		}
		send(log)
		time.Sleep(time.Second * time.Duration(rand.Intn(3)+1))
	}
}

func send(log map[string]interface{}) {
	body, _ := json.Marshal(log)
	http.Post("http://localhost:8001/ingest", "application/json", bytes.NewBuffer(body))
}

func randomLevel() string {
	return []string{"info", "warn", "error"}[rand.Intn(3)]
}
