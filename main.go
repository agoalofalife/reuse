package main

import (
	"fmt"
	"github.com/agoalofalife/reuse/commands"
	store "github.com/agoalofalife/storekeeper"
	"github.com/urfave/cli"
	"os"
	"time"
)

type Application struct {
	Container *store.Store
}

type manager struct {
	command []cli.Command
}

func (manager *manager) add(command cli.Command) {
	manager.command = append(manager.command, command)
}

func main() {
	app := Application{store.New()}
	app.Container.SetInstance(`command`, cli.NewApp())

	appCommand := app.Container.Extract(`command`).(*cli.App)

	manager := manager{}
	manager.add(commands.CopyConfig())
	manager.add(commands.StartServer())

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "version=%s\n", c.App.Version)
	}
	appCommand.Name = "reuse"
	appCommand.Authors = []cli.Author{
		cli.Author{
			Name:  "Ilya Chubarov",
			Email: "agoalofalife@gmail.com",
		},
	}
	appCommand.Compiled = time.Now()
	appCommand.Usage = "Cli application"

	appCommand.Flags = []cli.Flag{
		cli.DurationFlag{Name: "howlong, H", Value: time.Second * 3},
	}

	appCommand.Commands = manager.command
	appCommand.Run(os.Args)
}
