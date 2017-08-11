package reuse

import (
	"github.com/astaxie/beego/config"
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
	router := app.Container.Extract(`router`).(*Router).responsible
	loger := app.Container.Extract(`log`).(*Log)

	router.PathPrefix("/" + s.StaticUrl).Handler(
		http.StripPrefix("/"+s.StaticUrl, http.FileServer(http.Dir(os.ExpandEnv("$GOPATH")+"/"+s.StaticPath))))
	// TODO fix ExpandEnv -> os.Getwd()
	loger.Log.Notice(`Server is running on port ` + s.Port + `...`)
	http.ListenAndServe(`:`+s.Port, router)
}
