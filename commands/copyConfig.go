package commands

import (
	"github.com/urfave/cli"
	"os"
	"fmt"
	"github.com/shiena/ansicolor"
	"github.com/agoalofalife/reuse/support"
)

const configName  = "example.config.json"
const goPath = "$GOPATH"

func CopyConfig() cli.Command {
	// copy config json from stubs to base dir $GOPATH
	 return cli.Command{
		 Name:  "copy-config",
		 Usage: "Adds a copy of the configuration file",
		 Action: func(c *cli.Context) error {
			 dir, _ := os.Getwd()
			 w := ansicolor.NewAnsiColorWriter(os.Stdout)
			 text := "%sFile %s" + os.ExpandEnv(goPath) + "/" + "reuse.config.json " +"%screate success \n"
			 fmt.Fprintf(w,text, "\x1b[32m", "\x1b[92m", "\x1b[32m")
			 support.Cp(os.ExpandEnv(goPath) + "/" + "reuse.config.json", dir + "/stubs/" + configName,)
			 return nil
		 }}
}

