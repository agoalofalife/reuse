package reuse

import (
	"fmt"
	"github.com/agoalofalife/reuse/log"
	store "github.com/agoalofalife/storekeeper"
	"github.com/astaxie/beego/config"
	"github.com/gorilla/mux"
	"os"
	"path"
	"path/filepath"
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
	//Config config.Configer
	//// default router from gorilla mux
	//Router *mux.Router
}

// so start App..
func Run() {
	app.Container.Extract(`server`).(*Server).Run()
}

func Bootstrapping() {
	app = Application{store.New()}

	//fmt.Println(filepath.Abs(filepath.Dir(os.Args[0])))
	//workPath,_ :=  os.Getwd()
	fmt.Println(os.Getwd())
	AppPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	//appConfigPath := filepath.Join(workPath, "conf", "app.conf")
	//fmt.Println(appConfigPath)
	appConfigPath := filepath.Join(AppPath, "conf", "app.conf")
	fmt.Println(appConfigPath)
	os.Exit(0)

	//if !utils.FileExists(appConfigPath) {
	//	appConfigPath = filepath.Join(AppPath, "conf", "app.conf")
	//	//if !utils.FileExists(appConfigPath) {
	//	//	AppConfig = &beegoAppConfig{innerConfig: config.NewFakeConfig()}
	//	//	return
	//	//}
	//}
	//if err = parseConfig(appConfigPath); err != nil {
	//	panic(err)
	//}
	// TODO add the ability change path configuration for remote
	_, filename, _, _ := runtime.Caller(1)
	fileConf := path.Join(path.Dir(filename), relativePathConf)
	configuration, errConfig := config.NewConfig(typeConf, fileConf)
	fmt.Println(filename, `???a`)
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
