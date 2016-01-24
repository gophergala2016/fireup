package main

import (
	"github.com/codegangsta/cli"
)

func Serve(c *cli.Context) {
	host := c.String("o")
	port := c.Int("p")
	go Dispatcher()
	go StartPushListener(host, port+1)
	StartHttpListener(host, port)
}

func Push(c *cli.Context) {
	host := c.String("s")
	port := c.Int("p")
	args := c.Args()

	if len(args) > 0 {
		PushFile(host, port, args[0])
	} else {
		PushFile(host, port, "")
	}

}
