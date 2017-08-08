package reuse

import (
	"github.com/agoalofalife/reuse/log"
	store "github.com/agoalofalife/storekeeper"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
)

var app Application

type Application struct {
	Container *store.Store
}

type manager struct {
	command []cli.Command
}

func (manager *manager) add(command cli.Command) {
	manager.command = append(manager.command, command)
}

func Run() {
	app = Application{store.New()}
	loadDependencies()
	app.Container.Extract(`server`).(*Server).Run()
	//
	//app.Container.Bind(`core`, func(s *store.Store) *core.Server {
	//	return core.NewServer(os.ExpandEnv(goPath) + "/" + "reuse.config.json")
	//})
	//
	//appCommand := app.Container.Extract(`command`).(*cli.App)
	//
	//manager := manager{}
	//manager.add(commands.CopyConfig())
	//manager.add(commands.StartServer())
	//
	//cli.VersionPrinter = func(c *cli.Context) {
	//	fmt.Fprintf(c.App.Writer, "version=%s\n", c.App.Version)
	//}
	//appCommand.Name = "reuse"
	//appCommand.Authors = []cli.Author{
	//	cli.Author{
	//		Name:  "Ilya Chubarov",
	//		Email: "agoalofalife@gmail.com",
	//	},
	//}
	//appCommand.Compiled = time.Now()
	//appCommand.Usage = "Cli application"
	//
	//appCommand.Flags = []cli.Flag{
	//	cli.DurationFlag{Name: "howlong, H", Value: time.Second * 3},
	//}
	//
	//appCommand.Commands = manager.command
	//appCommand.Run(os.Args)
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
