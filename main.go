package main

import (
	"fmt"
	"github.com/agoalofalife/reuse/commands"
	"github.com/urfave/cli"
	"os"
	"time"
)

type manager struct {
	command []cli.Command
}

func (manager *manager) add(command cli.Command) {
	manager.command = append(manager.command, command)
}

func main() {
	app := cli.NewApp()
	manager := manager{}
	manager.add(commands.CopyConfig())
	manager.add(commands.StartServer())

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "version=%s\n", c.App.Version)
	}
	app.Name = "reuse"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Ilya Chubarov",
			Email: "agoalofalife@gmail.com",
		},
	}
	app.Compiled = time.Now()
	app.Usage = "Cli application"

	app.Flags = []cli.Flag{
		cli.DurationFlag{Name: "howlong, H", Value: time.Second * 3},
	}

	app.Commands = manager.command
	app.Run(os.Args)
}
