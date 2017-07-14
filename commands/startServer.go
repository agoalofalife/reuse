package commands

import (
	"github.com/urfave/cli"
	"github.com/agoalofalife/reuse/core"
	"os"
)

func StartServer() cli.Command {
	// run server start
	return cli.Command{
		Name:  "start",
		Usage: "Start server",
		Action: func(c *cli.Context) error {
			path := os.ExpandEnv(goPath) + "/" + "reuse.config.json"
			server := core.NewServer(path)
			server.Run()
			return nil
		}}
}
