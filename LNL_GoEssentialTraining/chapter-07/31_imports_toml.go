package main

import (
	"fmt"
	toml "github.com/pelletier/go-toml"
	"log"
	"os"
)

type ConfigHolder struct {
	Login struct {
		User string
		Password string
	}
}

func main() {
	fileHandler, err := os.Open("config.toml")

	if err != nil {
		log.Fatalf("Could not find a config file, err: %s", err)
	}
	defer fileHandler.Close()

	configDTO := &ConfigHolder{}
	tomlDecoder := toml.NewDecoder(fileHandler)

	if err := tomlDecoder.Decode(configDTO); err != nil {
		log.Fatalf("Can't decode toml file, %s", err)
	}

	fmt.Printf("%+v\n", configDTO)
}
