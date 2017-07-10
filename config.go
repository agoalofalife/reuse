package reuse

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Application struct {
	ServerConfig `json:"server"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

type Config struct {
	Application `json:"application"`
}

func (e Config) export(path string) *Config {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	config := new(Config)
	byteValue, _ := ioutil.ReadAll(file)

	json.Unmarshal(byteValue, &config)
	return config
}
