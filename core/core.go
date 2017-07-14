package core

import (
	"net/http"
	"github.com/agoalofalife/reuse/config"
	"github.com/agoalofalife/reuse/log"
)

type Server struct {
	config *config.Config
	log *log.Log
}

// create server
func NewServer(path string) *Server {
	config := config.Config{}
	configPointer := config.Export(path)
	return &Server{configPointer, log.NewLog()}
}

// initialization server
func (server Server) Run() {
	server.log.Log.Notice(`Server is running on port ` + server.config.Port + `...`)
	http.ListenAndServe(`:` + server.config.Port, nil)
}
