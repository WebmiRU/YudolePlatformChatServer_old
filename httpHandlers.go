package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("Hello world!"))
}

func modulesIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	model := GetModulesResponse{modules}
	response, _ := json.Marshal(model)

	w.Write(response)
}

func modulesIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id := mux.Vars(r)["id"]

	if _, ok := modules[id]; !ok {
		w.WriteHeader(404)
		return
	}

	model := GetModulesIdResponse{
		Payload: modules[id],
	}

	response, _ := json.Marshal(model)

	w.Write(response)
}

func modulesIdSetAutostartHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow", "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS")

	id := mux.Vars(r)["id"]
	state := mux.Vars(r)["state"]

	if _, ok := modules[id]; !ok {
		w.WriteHeader(404)
		return
	}

	switch r.Method {
	case "PUT":
		model := GetModulesResponse{modules}
		response, _ := json.Marshal(model)

		if state == "1" && modules[id].Autostart {
			fmt.Println("11")
		} else if state == "1" && !modules[id].Autostart {
			fmt.Println("10")
		} else if state == "0" && modules[id].Autostart {
			fmt.Println("01")
		} else if state == "0" && !modules[id].Autostart {
			fmt.Println("00")
		}

		w.Write(response)

	case "OPTIONS":
		w.WriteHeader(200)

	default:
		w.WriteHeader(404)
	}
}
