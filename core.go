package reuse

import (
	"net/http"
)

type Server struct {
	*Config
	*Log
}

// create server
func NewServer(path string) *Server {
	config := Config{}
	configPointer := config.export(path)
	return &Server{configPointer, newLog()}
}
// initialization server
func (server Server) Run() {
	server.log.Notice(`Server is running on port ` + server.Config.Port + `...`)
	http.ListenAndServe(`:` + server.Config.Port, nil)
}
