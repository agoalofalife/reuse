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
	//[]cli.Command{
	//{
	//	Name:    "generate-config",
	//	Usage:   "Adds a copy of the configuration file",
	//	Action:  func(c *cli.Context) error {
	//		fmt.Println("added task: ", c.Args().First())
	//		return nil
	//	},
	//},
	//{
	//	Name:        "template",
	//	Usage:       "options for task templates",
	//	Subcommands: []cli.Command{
	//		{
	//			Name:  "add",
	//			Usage: "add a new template",
	//			Action: func(c *cli.Context) error {
	//				fmt.Println("new task template: ", c.Args().First())
	//				return nil
	//			},
	//		},
	//		{
	//			Name:  "remove",
	//			Usage: "remove an existing template",
	//			Action: func(c *cli.Context) error {
	//				fmt.Println("removed task template: ", c.Args().First())
	//				return nil
	//			},
	//		},
	//	},
	//},
	//}
	app.Run(os.Args)
}
