package reuse

import (
	"net/http"
	"log"
)

type Server struct {
	*Config
}

// create server
func NewServer(path string) *Server {
	config := Config{}
	configPointer := config.export(path)
	return &Server{configPointer}
}
// initialization server
func (server Server) Run() {
	log.Println(`Server is running on port ` + server.Config.Port)
	http.ListenAndServe(`:` + server.Config.Port, nil)
}
