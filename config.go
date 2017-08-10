package reuse

import (
	"github.com/astaxie/beego/config"
)

const (
	typeConf = `ini`
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

func (c Config) LoadModule(app Application) bool {
	configuration, err := config.NewConfig(typeConf, c.pathToFile)

	if err != nil {
		panic(err.Error())
	}
	app.Container.SetInstance(`config`, configuration)
	return true
}
