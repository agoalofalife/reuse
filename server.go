package reuse

import (
	"github.com/agoalofalife/reuse/log"
	"github.com/astaxie/beego/config"
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

func NewServer() (server *Server) {
	conf := app.Container.Extract(`config`).(config.Configer)

	if conf.String(`port`) != "" {
		server := &Server{conf.String(`port`), conf.String(`staticPath`), conf.String(`staticUrl`)}
		return server
	}
	panic(`Not set port for server configuration!`)
}

func (s *Server) LoadModule(app Application) bool {
	app.Container.SetInstance(`server`, s)
	return true
}

func (s Server) Run() {
	// clean first symbol '/'
	if strings.Index(s.StaticUrl, "/") == 0 {
		s.StaticUrl = s.StaticUrl[1:]
	}
	router := app.Container.Extract(`router`).(*mux.Router)
	loger := app.Container.Extract(`log`).(*log.Log)

	router.PathPrefix("/" + s.StaticUrl).Handler(
		http.StripPrefix("/"+s.StaticUrl, http.FileServer(http.Dir(os.ExpandEnv("$GOPATH")+"/"+s.StaticPath))))
	loger.Log.Notice(`Server is running on port ` + s.Port + `...`)
	http.ListenAndServe(`:`+s.Port, router)
}
