package reuse

import (
	"github.com/agoalofalife/reuse/log"
	store "github.com/agoalofalife/storekeeper"
	"github.com/gorilla/mux"
)

var app Application

type Application struct {
	Container *store.Store
}

func Run() {
	app = Application{store.New()}
	loadDependencies()
	app.Container.Extract(`server`).(*Server).Run()
}

func loadDependencies() {
	server, err := NewServer(map[string]string{"port": "8080", "staticPath": "/", "staticUrl": "/static"})

	if err != "" {
		panic(err)
	}

	r := mux.NewRouter()
	app.Container.SetInstance(`server`, server)
	app.Container.SetInstance(`router`, r)
	app.Container.SetInstance(`log`, log.NewLog())

}
