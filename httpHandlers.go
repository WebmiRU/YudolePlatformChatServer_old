package main

import (
	"YudoleChatServer/packages/response"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func setCorsJsonHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Allow", "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-type")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("Hello world!"))
}

func modulesIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	model := response.GetModules{modules}
	resp, _ := json.Marshal(model)

	w.Write(resp)
}

func modulesIdHandler(w http.ResponseWriter, r *http.Request) {
	setCorsJsonHeaders(&w)

	id := mux.Vars(r)["id"]

	if _, ok := modules[id]; !ok {
		w.WriteHeader(404)
		return
	}

	model := response.GetModulesId{
		Payload: modules[id],
	}

	resp, _ := json.Marshal(model)

	w.Write(resp)
}

func modulesIdSetAutostartHandler(w http.ResponseWriter, r *http.Request) {
	setCorsJsonHeaders(&w)

	id := mux.Vars(r)["id"]
	state := mux.Vars(r)["state"]

	if _, ok := modules[id]; !ok {
		w.WriteHeader(404)
		return
	}

	switch r.Method {
	case "PUT":
		model := response.GetModules{Payload: modules}
		resp, _ := json.Marshal(model)

		if state == "1" && modules[id].Autostart {
			fmt.Println("11")
		} else if state == "1" && !modules[id].Autostart {
			fmt.Println("10")
		} else if state == "0" && modules[id].Autostart {
			fmt.Println("01")
		} else if state == "0" && !modules[id].Autostart {
			fmt.Println("00")
		}

		w.Write(resp)

	case "OPTIONS":
		w.WriteHeader(200)

	default:
		w.WriteHeader(404)
	}
}

func modulesIdStartHandler(w http.ResponseWriter, r *http.Request) {
	setCorsJsonHeaders(&w)

	id := mux.Vars(r)["id"]

	if _, ok := modules[id]; !ok {
		w.WriteHeader(404)
		return
	}

	switch r.Method {
	case "POST":
		modules[id].Start()

		model := response.GetModules{Payload: modules}
		resp, _ := json.Marshal(model)

		w.Write(resp)

	case "OPTIONS":
		w.WriteHeader(200)

	default:
		w.WriteHeader(404)
	}
}

func modulesIdStopHandler(w http.ResponseWriter, r *http.Request) {
	setCorsJsonHeaders(&w)

	id := mux.Vars(r)["id"]

	if _, ok := modules[id]; !ok {
		w.WriteHeader(404)
		return
	}

	switch r.Method {
	case "POST":
		modules[id].Stop()

		model := response.GetModules{Payload: modules}
		resp, _ := json.Marshal(model)

		w.Write(resp)

	case "OPTIONS":
		w.WriteHeader(200)

	default:
		w.WriteHeader(404)
	}
}
