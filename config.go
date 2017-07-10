package reuse

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Application `json:"application"`
}

type Application struct {
	Server `json:"server"`
}

type Server struct {
	Port int16 `json:"port"`
}

func (e Config) export(path string) *Config {
	file, _ := os.Open(path)
	defer file.Close()

	config  := new(Config)
	byteValue, _ := ioutil.ReadAll(file)

	json.Unmarshal(byteValue, &config)
	return  config
}