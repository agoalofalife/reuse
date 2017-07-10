package reuse

import (
	"net/http"
	"github.com/op/go-logging"
	"os"
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
	var lg = logging.MustGetLogger("example")
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	//backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled := logging.NewBackendFormatter(backend1, format)
	//backend1Leveled.SetLevel(logging.ERROR, "")
	logging.SetBackend(backend1Leveled)
	lg.Notice(`Server is running on port ` + server.Config.Port + `...`)
	//log.Println(`Server is running on port ` + server.Config.Port + `...`)
	http.ListenAndServe(`:` + server.Config.Port, nil)
}
