package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")

	for {
		fmt.Fprintf(w, "data: %d\n\n", rand.Intn(999))
		w.(http.Flusher).Flush()
		time.Sleep(1 * time.Second)
	}
}
