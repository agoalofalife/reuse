package core

import (
	"github.com/agoalofalife/reuse/config"
	"github.com/agoalofalife/reuse/log"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
)

type Router interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	PathPrefix(tpl string) *mux.Route
}

type Server struct {
	config *config.Config
	log    *log.Log
	router *mux.Router
}

// create server
func NewServer(path string) *Server {

	conf := config.Config{}
	configPointer := conf.Export(path)
	return &Server{configPointer, log.NewLog(), mux.NewRouter()}
}

// initialization server
func (server Server) Run() {
	// clean first symbol '/'
	if strings.Index(server.config.StaticUrl, "/") == 0 {
		server.config.StaticUrl = server.config.StaticUrl[1:]
	}

	server.router.PathPrefix("/" + server.config.StaticUrl).Handler(
		http.StripPrefix("/"+server.config.StaticUrl, http.FileServer(http.Dir(os.ExpandEnv("$GOPATH")+"/"+server.config.StaticPath))))
	server.log.Log.Notice(`Server is running on port ` + server.config.Port + `...`)
	http.ListenAndServe(`:`+server.config.Port, server.router)
}
