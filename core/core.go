package core

import (
	"net/http"
	"github.com/agoalofalife/reuse/cli"
	"github.com/agoalofalife/reuse/config"
	"github.com/agoalofalife/reuse/log"
)

type Server struct {
	config *config.Config
	log *log.Log
	command *cli.Command
}

// create server
func NewServer(path string) *Server {
	config := config.Config{}
	configPointer := config.Export(path)
	cli := &cli.Command{}
	cli.Start()
	return &Server{configPointer, log.NewLog(), cli}
}

// initialization server
func (server Server) Run() {

	server.log.Notice(`Server is running on port ` + server.config.Port + `...`)
	http.ListenAndServe(`:` + server.config.Port, nil)
}
