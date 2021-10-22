package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	dataaccess "handler/joiner-getby/func/dataaccess"

	"github.com/joho/godotenv"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file -tags debug")
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
