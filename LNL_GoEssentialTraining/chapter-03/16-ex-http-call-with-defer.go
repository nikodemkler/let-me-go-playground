package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		return "", fmt.Errorf("Wrong")
	}

	return contentType, nil
}

func main() {
	ctype, err := contentType("https://www.linkeasddin.org")

	if err != nil {
		fmt.Println("-----------")
		fmt.Printf("Err: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}