package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	db = map[string]interface{}{}
	dbLock sync.Mutex
)

type Entry struct {
	Key string			`json:"key"`
	Value interface{}	`json:"value"`
}

func sendResponse(entry *Entry, wrt http.ResponseWriter) {
	wrt.Header().Set("Content-Type", "application/json")
	respEncoder := json.NewEncoder(wrt)

	if err := respEncoder.Encode(entry); err != nil {
		http.Error(wrt, err.Error(), http.StatusInternalServerError)
	}
}

func kvPostHandler(wrt http.ResponseWriter, req *http.Request) {
	//	Decode request
	defer req.Body.Close()

	entry := &Entry{}
	requestDecoder := json.NewDecoder(req.Body)
	if err := requestDecoder.Decode(entry); err != nil {
		http.Error(wrt, err.Error(), http.StatusBadRequest)
	}

	// db lock init
	dbLock.Lock()
	defer dbLock.Unlock()

	db[entry.Key] = entry.Value

	// sens the current result as the response
	sendResponse(entry, wrt)
}

func kvGetHandler(wrt http.ResponseWriter, req *http.Request) {
	key := req.URL.Path[7:]

	dbLock.Lock()
	defer dbLock.Unlock()

	value, ok := db[key]
	if !ok {
		http.Error(wrt, fmt.Sprintf("Key %q", key), http.StatusNotFound)
	}

	respEntry := &Entry{
		Key: key,
		Value: value,
	}

	sendResponse(respEntry, wrt)
}

func main() {
	http.HandleFunc("/store", kvPostHandler)
	http.HandleFunc("/store/", kvGetHandler)
	if err := http.ListenAndServe(":8099", nil); err != nil {
		log.Fatalf("Error %s", err)
	}
}