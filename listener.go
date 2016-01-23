package main

import (
	"log"
	"fmt"
	"net"
	"bufio"
	"strings"
	"path"
	"io/ioutil"
)

func getFileType(fname string) string {
	return path.Ext(fname)
}

func handlePush(conn net.Conn) {
	r := bufio.NewReader(conn)

	header, err := r.ReadString('\n')

	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fileType := getFileType(strings.TrimSpace(header))
	log.Println("recvd:", fileType)

	data, err := ioutil.ReadAll(conn)

	

	conn.Close()

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
