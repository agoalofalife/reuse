package core

import (
	"net/http"
	"github.com/agoalofalife/reuse/config"
	"github.com/agoalofalife/reuse/log"
	"os"
	"github.com/gorilla/mux"
)

type Server struct {
	config *config.Config
	log *log.Log
	router *mux.Router
}


// create server
func NewServer(path string) *Server {
	config := config.Config{}
	configPointer := config.Export(path)
	return &Server{configPointer, log.NewLog(), mux.NewRouter()}
}

// initialization server
func (server Server) Run() {
	//fs := http.FileServer(http.Dir(os.ExpandEnv("$GOPATH") + "/" + server.config.StaticUrl))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	server.router.PathPrefix("/" + server.config.StaticUrl).Handler(
		http.StripPrefix("/" + server.config.StaticUrl, http.FileServer(http.Dir(os.ExpandEnv("$GOPATH") + "/" + server.config.StaticPath))))
	server.log.Log.Notice(`Server is running on port ` + server.config.Port + `...`)
	http.ListenAndServe(`:` + server.config.Port, server.router)
}
