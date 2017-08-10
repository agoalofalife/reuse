package reuse

import (
	"github.com/agoalofalife/reuse/supports/files"
	store "github.com/agoalofalife/storekeeper"
	"os"
	"path/filepath"
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

	NewConfig(appConfigPath).LoadModule(app)
	NewServer().LoadModule(app)
	NewLog().LoadModule(app)
	NewRouter().LoadModule(app)

}

func kernel() Application {
	if app.Container == nil {
		app.Container = store.New()
	}
	return app
}
