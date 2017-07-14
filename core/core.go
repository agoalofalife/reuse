package core

import (
	"net/http"
	"github.com/agoalofalife/reuse/config"
	"github.com/agoalofalife/reuse/log"
	"os"
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
	fs := http.FileServer(http.Dir(os.ExpandEnv("$GOPATH") + "/" + server.config.StaticUrl))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	server.log.Log.Notice(`Server is running on port ` + server.config.Port + `...`)
	http.ListenAndServe(`:` + server.config.Port, nil)
}
