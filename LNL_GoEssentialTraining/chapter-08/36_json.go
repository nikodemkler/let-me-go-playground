package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var jsonData = `{
	"user": "Foo",
	"type": "deposit",
	"amount": 100.3
}`

type Request struct {
	Login string 	`json:"user"`
	Type string		`json:"type"`
	Amount float64 	`json:"amount"`
}

func main() {
	jsonDataHandler := bytes.NewBufferString(jsonData) // simulate a file / socket

	// decode request
	dataReader := json.NewDecoder(jsonDataHandler)

	request := &Request{}
	if err := dataReader.Decode(request); err != nil {
		log.Fatalf("Error during decoding, %s", err)
	}

	fmt.Printf("Request: %+v\n", request)

	// create response
	prevBalance := 850.0 // loaded from DB
	resp := map[string]interface {} {
		"ok": true,
		"balance": prevBalance + request.Amount,
	}

	// encode it back to JSON
    responseEncoder := json.NewEncoder(os.Stdout)

    if err := responseEncoder.Encode(resp); err != nil {
    	log.Fatalf("error: can't encode a response, %s", err)
	}
}