package main

import (
	"github.com/codegangsta/cli"
	"fmt"
	"os"
)

func Serve(c *cli.Context) {
	host := c.String("o")
	port := c.Int("p")
	go StartPushListener(host, port + 1)
	StartHttpListener(host, port)
}

func Push(c *cli.Context) {
	host := c.String("s")
	port := c.Int("p")
	args := c.Args()

	var err error
	if len(args) > 0 {
		err = PushFile(host, port, args[0])
	} else {
		err = PushFile(host, port, "")
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	}
}

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
