package reuse

import (
	"github.com/astaxie/beego/config"
	"github.com/jinzhu/gorm"
)

type DataBase struct {
	Name     string
	Login    string
	Password string
	Type     string
}

func InstallDatabase(dialect string, parameters string) {
	var (
		db  *gorm.DB
		err error
	)
	configuration := app.Container.Extract(`config`).(config.Configer)
	if parameters != "" {
		db, err = gorm.Open(
			configuration.String(`typeDatabase`),
			configuration.String(`loginDatabase`)+
				":"+
				configuration.String(`passwordDatabase`)+
				"@/"+
				configuration.String(`nameDatabase`)+"?charset=utf8&parseTime=True&loc=Local")
	} else {
		db, err = gorm.Open(dialect, parameters)
	}

	if err != nil {
		panic("failed to connect database")
	}

	app.Container.SetInstance(`database`, db)
}
