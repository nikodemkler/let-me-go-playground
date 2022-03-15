package main

import (
	"fmt"
	"net/http"
)

func returnContentType(url string, out chan string) string {
	resp, err := http.Get(url)

	if err != nil {
		out <- fmt.Sprintf("error %+v", err)
		return ""
	}

	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	out <- contentType

	return contentType
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	resultChannel := make(chan string)
	for _, url := range urls {
		go func(url string) {
			contentType := returnContentType(url, resultChannel)
			fmt.Printf("Print: %s\n", contentType)
		}(url)
	}

	for range urls {
		outMessage := <-resultChannel
		fmt.Printf("Out: %s\n", outMessage)
	}
}