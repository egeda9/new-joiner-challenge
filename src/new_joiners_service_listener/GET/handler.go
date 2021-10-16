package main

import (
	"encoding/json"
	"log"
	"net/http"

	dataaccess "handler/joiner-get/func/dataaccess"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	joiners, err := dataaccess.Get()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(joiners) <= 0 {
		http.Error(w, "No Results", http.StatusNoContent)
		return
	}

	jsonResp, err := json.Marshal(joiners)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/Joiner", requestHandler)

	err := http.ListenAndServe(":5006", mux)
	log.Fatal(err)
}
