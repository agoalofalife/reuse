package commands

import (
	"github.com/urfave/cli"
	"os"
	"io"
	"fmt"
)

const configName  = "example.config.json"
const gopath = "$GOPATH"

func CopyConfig() cli.Command {
	 return cli.Command{
		 Name:  "copy-config",
		 Usage: "Adds a copy of the configuration file",
		 Action: func(c *cli.Context) error {
			 dir, _ := os.Getwd()

			 fmt.Println("added task: ", dir + "/stubs/" + configName)
			 cp(os.ExpandEnv(gopath) + "/" + "reuse.config.json", dir + "/stubs/" + configName,)
			 //fmt.Println("added task: ",dir + "/" + configName)
			 //fmt.Println("added task: ", c.Args().First())
			 return nil
		 }}
}

func cp(dst, src string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	// no need to check errors on read only file, we already got everything
	// we need from the filesystem, so nothing can go wrong now.
	defer s.Close()
	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}