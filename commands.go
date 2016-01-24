package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func Serve(c *cli.Context) {
	host := c.String("o")
	port := c.Int("p")
	go StartPushListener(host, port+1)
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
