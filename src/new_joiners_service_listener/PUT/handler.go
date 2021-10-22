package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	dataaccess "handler/joiner-update/func/dataaccess"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {
		j := new(dataaccess.Joiner)
		err := json.NewDecoder(r.Body).Decode(&j)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

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

		existingJoiner, err := dataaccess.Get(joinerId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		j.UpdateJoiner(existingJoiner.Id)

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
}

// We could fix the port number, but cloud environments normally require
// some flexibility on defining the server port. This is how it would work
// in Azure.
func getHTTPPort() int {
	httpPort := 8080
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		httpPort, err := strconv.Atoi(val)
		if err == nil {
			return httpPort
		}
	}
	return httpPort
}

func main() {
	httpPort := getHTTPPort()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/Joiner", requestHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux)
	log.Fatal(err)
}
