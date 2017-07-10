package reuse

import (
	"os"
	"encoding/json"
	"log"
	"fmt"
)
const (
	error_json_file = "Not work decode file json"
	error_file_not_exist = "File not exist"
)

type Config struct {
	Applications Application
}

type Application struct {
	Server
}

type Server struct {
	port string `json:"port"`
}
func (e Config) Export(path string) *Config {
	fileExist(path)
	file, _ := os.Open(path)

	decoder := json.NewDecoder(file)
	log.Println(decoder)
	os.Exit(2)
	config  := new(Config)
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal(error_json_file)
	}
	return  config
}




func fileExist(path string) (is bool, err error)  {
	if _, err := os.Stat(path); os.IsNotExist(err) {
	return false, fmt.Errorf("%s", error_file_not_exist)
	}
	return true, nil
}