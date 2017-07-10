package cli

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)
type Command struct{

}

func (command *Command) Start()  {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		fmt.Printf("Hello %q", c.Args().Get(0))
		return nil
	}

	app.Run(os.Args)
}
