package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"sync"
)

type sseClient struct {
	W    *http.ResponseWriter
	R    *http.Request
	Chan *chan any
}

func (c *sseClient) Send(message any) error {
	*c.Chan <- message
	return nil
}

func (c *sseClient) Drop() error {
	sseClientsMutex.Lock()
	idx := slices.Index(sseClients, c)
	sseClients = slices.Delete(sseClients, idx, idx+1)
	sseClientsMutex.Unlock()

	sseEventSubsMutex.Lock()
	for event, clients := range sseEventSubs {
		idx := slices.Index(clients, c)

		if idx == -1 {
			continue
		}

		sseEventSubs[event] = slices.Delete(clients, idx, idx+1)
	}
	sseEventSubsMutex.Unlock()

	return nil
}

var sseClientsMutex sync.Mutex
var sseEventSubsMutex sync.Mutex

// var sseClients = make(map[*http.ResponseWriter]*sseClient)
var sseClients = make([]*sseClient, 0)
var sseEventSubs = make(map[string][]*sseClient)
var sseChan = make(chan Message)

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")

	ch := make(chan any)
	client := &sseClient{
		W:    &w,
		R:    r,
		Chan: &ch,
	}

	subscribe := r.URL.Query()["subscribe[]"]

	sseEventSubsMutex.Lock()
	for _, event := range subscribe {
		if slices.Contains(events, event) {
			sseEventSubs[event] = append(sseEventSubs[event], client)
		} else {
			log.Println("Unknown event type:", event)
		}
	}
	sseEventSubsMutex.Unlock()

	sseClientsMutex.Lock()
	sseClients = append(sseClients, client)
	sseClientsMutex.Unlock()

loop:
	for {
		select {
		case message := <-ch:
			msg, _ := json.Marshal(message)
			if _, err := fmt.Fprintf(w, "data: %s\n\n", msg); err != nil {
				break loop
			}

			w.(http.Flusher).Flush()

		case <-r.Context().Done():
			break loop
		}
	}

	client.Drop()
	fmt.Println("SSE CLIENT DISCONNECTED")
}
