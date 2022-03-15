package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Name string			`json:"name"`
	PublicRepos int		`json:"public_repos"`
}

func userInfo(login string) (*User, error) {
	var githubEndpoint string = fmt.Sprintf("https://api.github.com/users/%s", login)

	resp, err := http.Get(githubEndpoint)

	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("error: %s", err)
		}
	}(resp.Body)

	user := &User{}
	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func main() {
	user, err := userInfo("nikodemkler")
	if err != nil {
		log.Fatalf("Can't fetch the user, %s", err)
	}

	fmt.Println(user)
}