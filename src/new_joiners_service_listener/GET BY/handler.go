package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	dataaccess "handler/joiner-getby/func/dataaccess"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	id := query["id"]

	if (len(id) <= 0) || (id[0] == "") || (id[0] == "0") {
		http.Error(w, "Missing/Invalid id query string parameter", http.StatusBadRequest)
		return
	}

	joinerId, err := strconv.Atoi(id[0])

	if err != nil {
		http.Error(w, "Invalid id query string parameter", http.StatusBadRequest)
		return
	}

	joiner, err := dataaccess.Get(joinerId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonResp, err := json.Marshal(joiner)

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

	err := http.ListenAndServe(":5008", mux)
	log.Fatal(err)
}
