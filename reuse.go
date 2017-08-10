package reuse

import (
	"github.com/agoalofalife/reuse/log"
	"github.com/agoalofalife/reuse/supports/files"
	store "github.com/agoalofalife/storekeeper"
	"github.com/astaxie/beego/config"
	"github.com/gorilla/mux"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var app Application

type Application struct {
	// container store struct in application
	Container *store.Store
}

// so start App..
func Run() {
	bootstrapping()
	app.Container.Extract(`server`).(*Server).Run()
}

func bootstrapping() {
	app = kernel()

	//AppPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//appConfigPath = filepath.Join(AppPath, "conf", "app.conf")
	workDir, _ := os.Getwd()
	appConfigPath := filepath.Join(workDir, "conf", "app.conf")

	if !files.FileExists(appConfigPath) {
		panic(`File configuration is not set!`)
	}

	NewConfig(appConfigPath)

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
}

func kernel() Application {
	if app.Container == nil {
		app.Container = store.New()
	}
	return app
}
