package reuse

import (
	"github.com/agoalofalife/reuse/log"
	store "github.com/agoalofalife/storekeeper"
	"github.com/astaxie/beego/config"
	"github.com/gorilla/mux"
	"path"
	"runtime"
)

const (
	relativePathConf = `/config/app.conf`
	typeConf         = `ini`
)

var app Application

type Application struct {
	// container store struct in application
	Container *store.Store
	// configuration module from beego
	Config config.Configer
}

// so start App..
func Run() {
	loadDependencies()
	app.Container.Extract(`server`).(*Server).Run()
}

func loadDependencies() {
	// TODO add the ability change path configuration for remote
	_, filename, _, _ := runtime.Caller(1)
	fileConf := path.Join(path.Dir(filename), relativePathConf)
	configuration, errConfig := config.NewConfig(typeConf, fileConf)

	server, err := NewServer(map[string]string{"port": configuration.String(`port`),
		"staticPath": configuration.String(`staticPath`),
		"staticUrl":  configuration.String(`staticUrl`)})

	if err != nil {
		panic(err.Error())
	}

	if errConfig != nil {
		panic(errConfig.Error())
	}

	r := mux.NewRouter()
	app.Container.SetInstance(`config`, configuration)
	app.Container.SetInstance(`server`, server)
	app.Container.SetInstance(`router`, r)
	app.Container.SetInstance(`log`, log.NewLog())

	app = Application{store.New(), configuration}
}
