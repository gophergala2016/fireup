package main

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/websocket"
)

func StartHttpListener(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("starting http listener on http://%s\n", addr)

	http.Handle("/ws", websocket.Handler(BroadcastHandler))
	http.HandleFunc("/slides/", SlideHandler)
	http.HandleFunc("/assets/", AssetHandler)
	http.HandleFunc("/", RootHandler)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Println("ERROR:", err)
	}
}
