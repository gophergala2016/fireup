package main

import (
	"log"
	"fmt"
)

func StartPushListener(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Println("starting push listener on", addr)
}
