package reuse

import (
	"github.com/astaxie/beego/config"
	"github.com/agoalofalife/storekeeper"
)

const (
	relativePathConf = `/config/app.conf`
	typeConf         = `ini`
)

type Config struct {
	// path to file
	pathToFile string
	// responsible
	responsible config.Configer
}

// create new Structure Config
func NewConfig(configPath string) *Config {
	configuration, errConfig := config.NewConfig(typeConf, configPath)
	if errConfig != nil {
		panic(errConfig.Error())
	}
	pointerConfig := &Config{configPath, configuration}
	app.Container.SetInstance(`config`, pointerConfig)
	return pointerConfig
}

func LoadModule(app Application) bool {

}