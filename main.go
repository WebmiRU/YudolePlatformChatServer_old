package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"slices"
	"strings"
)

var config Config
var currentDir string
var modules = make(map[string]Module)

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
	extList, _ := os.ReadDir("./modules")

	for _, ext := range extList {
		extDir := "./modules/" + ext.Name()
		configBytes, _ := os.ReadFile(extDir + string(os.PathSeparator) + "/config.json")

		var module Module
		json.Unmarshal(configBytes, &module)

		if len(module.Command) >= 2 && module.Command[0:2] == "./" {
			module.Command = currentDir + strings.Replace(module.Command, "./", string(os.PathSeparator), 1)
		}

		//module.Key = ext.Name()
		module.Path = currentDir + string(os.PathSeparator) + "modules" + string(os.PathSeparator) + ext.Name()
		module.IsActive = slices.Contains(config.ActiveModules, ext.Name())

		modules[ext.Name()] = module
	}

	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/api/modules", modulesHandler)
	router.HandleFunc("/api/modules/{id}", modulesIdHandler)
	http.Handle("/", router)

	http.ListenAndServe(":80", nil)
}
