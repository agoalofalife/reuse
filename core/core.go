package core

import (
	"net/http"
	"github.com/agoalofalife/reuse/cli"
)

type Server struct {
	*Config
	*Log
	*cli.Command
}

// create server
func NewServer(path string) *Server {
	config := Config{}
	configPointer := config.export(path)
	cli := &cli.Command{}
	cli.Start()
	return &Server{configPointer, newLog(), cli}
}
// initialization server
func (server Server) Run() {
	server.log.Notice(`Server is running on port ` + server.Config.Port + `...`)
	http.ListenAndServe(`:` + server.Config.Port, nil)
}
