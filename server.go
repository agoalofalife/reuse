package reuse

import (
	"github.com/agoalofalife/reuse/log"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
)

type Server struct {
	Port       string
	StaticPath string
	StaticUrl  string
}

func NewServer(parameters map[string]string) (server *Server, err string) {
	if parameters["port"] != "" && parameters["staticPath"] != "" && parameters["staticUrl"] != "" {
		return &Server{parameters["port"], parameters["staticPath"], parameters["staticUrl"]}, err
	} else {
		err = `Check input parameters : port , staticPath, staticUrl`
		return nil, err
	}

}

func (server Server) Run() {
	// clean first symbol '/'
	if strings.Index(server.StaticUrl, "/") == 0 {
		server.StaticUrl = server.StaticUrl[1:]
	}
	router := app.Container.Extract(`router`).(*mux.Router)
	loger := app.Container.Extract(`log`).(*log.Log)

	router.PathPrefix("/" + server.StaticUrl).Handler(
		http.StripPrefix("/"+server.StaticUrl, http.FileServer(http.Dir(os.ExpandEnv("$GOPATH")+"/"+server.StaticPath))))
	loger.Log.Notice(`Server is running on port ` + server.Port + `...`)
	http.ListenAndServe(`:`+server.Port, router)
}
