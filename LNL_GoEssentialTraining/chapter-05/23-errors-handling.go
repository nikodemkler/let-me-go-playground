package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

type Config struct {

}

func readConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.Wrap(err, "Can't open config file")
	}

	defer file.Close()

	cfg := &Config{}
	return cfg, nil
}

func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func main() {
	setupLogging()

	cfg, err := readConfig("/etc/fake.file")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		log.Printf("Error: %+v", err)
		os.Exit(1)
	}

	fmt.Println(cfg)
}