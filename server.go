package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func StartHttpListener(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("starting http listener on http://%s\n", addr)

	http.Handle("/ws", websocket.Handler(BroadcastHandler))
	http.HandleFunc("/slides/", SlideHandler)

	http.HandleFunc("/assets/css/", CssHandler)
	http.HandleFunc("/assets/js/", JsHandler)
	http.HandleFunc("/assets/fonts/", FontHandler)

	http.HandleFunc("/", RootHandler)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Println("ERROR:", err)
	}
}
