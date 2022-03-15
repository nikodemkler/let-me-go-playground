package main

import (
	"fmt"
	"sync"
)
import "net/http"

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		//returnType(url)
		wg.Add(1)
		go func(url string){
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}