package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"github.com/russross/blackfriday"
)

func handlePush(conn net.Conn) {
	data, err := ioutil.ReadAll(conn)

	if err != nil {
		log.Println("ERROR", err)
		return
	}

	output := blackfriday.MarkdownBasic(data)
	key := Save(string(output))

	chData <- key

}

func StartPushListener(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Println("starting push listener on", addr)

	ln, err := net.Listen("tcp", addr)

	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Println("ERROR:", err)
			return
		}
		go handlePush(conn)
	}
}
