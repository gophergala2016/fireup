package main

import (
	"golang.org/x/net/websocket"
)

type StringChan chan string

var (
	chData  = make(chan string)
	clients = make([]StringChan, 0)
)

func BroadcastHandler(ws *websocket.Conn) {
	ch := make(chan string)
	clients = append(clients, ch)
	for {
		key := <-ch // this is the key
		websocket.JSON.Send(ws, key)
	}
}
