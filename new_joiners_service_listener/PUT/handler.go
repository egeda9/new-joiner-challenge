package main

import (
	"encoding/json"
	"log"
	"net/http"

	dataaccess "handler/joiner-update/func/dataaccess"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	j := new(dataaccess.Joiner)
	err := json.NewDecoder(r.Body).Decode(&j)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	j.UpdateJoiner()

	response := make(map[string]string)
	response["message"] = "Joiner Updated"
	jsonResp, err := json.Marshal(response)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/NewJoinerUpdateFunction", requestHandler)

	err := http.ListenAndServe(":5005", mux)
	log.Fatal(err)
}
