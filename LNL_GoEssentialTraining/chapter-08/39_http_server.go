package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MathRequest struct {
	Op string		`json:"op"`
	Left float64	`json:"left"`
	Right float64	`json:"right"`
}

type MathResponse struct {
	Error string	`json:"error"`
	Result float64	`json:"result"`
}

func mathHandler(wrt http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)

	requestModel := &MathRequest{}
	if err := decoder.Decode(requestModel); err != nil {
		http.Error(wrt, err.Error(), http.StatusBadRequest)
		return
	}

	respModel := &MathResponse{}
	switch requestModel.Op {
	case "+":
		respModel.Result = requestModel.Left + requestModel.Right
	case "-":
		respModel.Result = requestModel.Left - requestModel.Right
	case "*":
		respModel.Result = requestModel.Left * requestModel.Right
	case "/":
		if requestModel.Right == 0.0 {
			respModel.Error = "division by zero"
		} else {
			respModel.Result = requestModel.Left / requestModel.Right
		}
	default:
		respModel.Error = fmt.Sprintf("unknown operation, %s", requestModel.Op)
	}

	wrt.Header().Set("Content-Type", "application/json")

	if respModel.Error != "" {
		wrt.WriteHeader(http.StatusBadRequest)
	}

	encoder := json.NewEncoder(wrt)

	if err := encoder.Encode(respModel); err != nil {
		http.Error(wrt, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/math", mathHandler)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}