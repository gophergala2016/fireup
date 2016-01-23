package main

import (
	"github.com/codegangsta/cli"
	"os"
)


func main() {

	app := cli.NewApp()
	app.Name = "fireup"
	app.Usage = "fireup stuff"

	app.Commands = []cli.Command{
		{
			Name: "serve", 
			Usage: "serve", 
			Aliases: []string{ "s" },
			Action: Serve, 
			Flags: []cli.Flag{
				cli.StringFlag{ Name: "o", Value: "0.0.0.0", Usage: "listen on HOST."},
				cli.IntFlag{ Name: "p", Value: 8080, Usage: "use PORT."},
			},
		},
		{
			Name: "push", 
			Usage: "push", 
			Aliases: []string{ "p" },
			Action: Push, 
			Flags: []cli.Flag{
				cli.StringFlag{ Name: "s", Value: "0.0.0.0", Usage: "remote server adderess."},
				cli.IntFlag{ Name: "p", Value: 8081, Usage: "remote server port."},
			},
		},
	}

	app.Run(os.Args)
}
