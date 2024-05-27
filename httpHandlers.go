package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("Hello world! 123"))
}

func modulesHandler(w http.ResponseWriter, r *http.Request) {
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

	if v, ok := modules[id]; ok {
		model := GetModulesIdResponse{
			Payload: v,
		}
		response, _ := json.Marshal(model)

		w.Write(response)
	} else {
		w.WriteHeader(404)
	}
}
