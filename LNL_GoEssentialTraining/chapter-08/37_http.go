package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Job struct {
	User string		`json:"user"`
	Action string	`json:"action"`
	Count int		`json:"count"`
}

func main() {

	response, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("Error in call, %s", err)
	}
	defer response.Body.Close()

	io.Copy(os.Stdout, response.Body)

	fmt.Println("------------")

	newJob := &Job{
		User: "John",
		Action: "punch",
		Count: 1,
	}

	var buff bytes.Buffer
	enc := json.NewEncoder(&buff)

	if err := enc.Encode(newJob); err != nil {
		log.Fatalf("Error during encoding, %s", err)
	}
	createResponse, err := http.Post("https://httpbin.org/post", "application/json", &buff)

	if err != nil {
		log.Fatalf("Error in creation, %s", err)
	}
	defer createResponse.Body.Close()

	io.Copy(os.Stdout, createResponse.Body)
}