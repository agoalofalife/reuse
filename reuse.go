package reuse

import (
	"github.com/agoalofalife/reuse/log"
	store "github.com/agoalofalife/storekeeper"
	"github.com/astaxie/beego/config"
	"github.com/gorilla/mux"
	"path"
	"runtime"
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
	if err != nil {
		panic(err.Error())
	}
	_, filename, _, _ := runtime.Caller(1)
	fileConf := path.Join(path.Dir(filename), "/config/app.conf")
	configuration, errConfig := config.NewConfig("ini", fileConf)

	if errConfig != nil {
		panic(errConfig.Error())
	}

	r := mux.NewRouter()
	app.Container.SetInstance(`config`, configuration)
	app.Container.SetInstance(`server`, server)
	app.Container.SetInstance(`router`, r)
	app.Container.SetInstance(`log`, log.NewLog())

}
