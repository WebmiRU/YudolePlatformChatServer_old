package main

import (
	"YudoleChatServer/packages/module"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	TCP uint8 = iota
	WS
	SSE
)

type IClient interface {
	Send(message any) error
}

var signals = make(chan os.Signal, 99)
var config Config
var currentDir string
var modules = make(map[string]*module.Module)
var eventSubs = make(map[string][]IClient)

func loadConfig() {
	configBytes, err := os.ReadFile("config.json")

	if err != nil {
		panic("Error while reading 'config.json' file")
	}

	if json.Unmarshal(configBytes, &config) != nil {
		panic("Error while parsing 'config.json' file")
	}

}

func main() {
	// Catch shutdown signals
	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-signals:
				for _, m := range modules {
					m.Stop()
				}

				log.Println("Shutting down...")
				os.Exit(0)
			}
		}

	}()

	loadConfig()
	go tcpServer()

	currentDir, _ = os.Getwd()
	moduleList, _ := os.ReadDir(currentDir + fmt.Sprintf("%c%s", os.PathSeparator, "modules"))

	for _, dir := range moduleList {
		path := currentDir + string(os.PathSeparator) + "modules" + string(os.PathSeparator) + dir.Name()

		var mod module.Module
		if err := mod.Load(path); err == nil {
			modules[dir.Name()] = &mod
		} else {
			log.Println(err)
		}

		if err := mod.Start(); err != nil {
			log.Println(err)
			return
		}
	}

	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/api/modules", modulesIndexHandler)
	router.HandleFunc("/api/modules/{id}", modulesIdHandler)
	router.HandleFunc("/api/modules/{id}/start", modulesIdStartHandler)
	router.HandleFunc("/api/modules/{id}/stop", modulesIdStopHandler)
	router.HandleFunc("/api/modules/{id}/autostart/{state:[0,1]}", modulesIdSetAutostartHandler)
	http.Handle("/", router)

	http.HandleFunc("/events", eventsHandler)

	http.ListenAndServe(":80", nil)
}
