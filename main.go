package main

import (
	"YudoleChatServer/packages/module"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var config Config
var currentDir string
var modules = make(map[string]*module.Module)

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
	loadConfig()

	currentDir, _ = os.Getwd()
	moduleList, _ := os.ReadDir(currentDir + fmt.Sprintf("%c%s", os.PathSeparator, "modules"))

	for _, dir := range moduleList {
		path := currentDir + "/modules/" + dir.Name()

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
	router.HandleFunc("/api/modules/{id}/autostart/{state:[0,1]}", modulesIdSetAutostartHandler)
	http.Handle("/", router)

	http.ListenAndServe(":80", nil)
}
