package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	RootTemplate.Execute(w, nil)
}

func SlideHandler(w http.ResponseWriter, r *http.Request) {
	key := path.Base(r.URL.Path)

	log.Println(key)
	fmt.Fprintf(w, Get(key))

}

func JsHandler(w http.ResponseWriter, r *http.Request) {
	asset := r.URL.Path[1:]

	data, err := Asset(asset)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	w.Header().Set("Content-Type", "application/javascript")

	w.Write(data)
}

func CssHandler(w http.ResponseWriter, r *http.Request) {

	asset := r.URL.Path[1:]

	data, err := Asset(asset)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	w.Header().Set("Content-Type", "text/css")

	w.Write(data)

}

func FontHandler(w http.ResponseWriter, r *http.Request) {

	asset := r.URL.Path[1:]

	data, err := Asset(asset)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	w.Header().Set("Content-Type", "font/opentype")

	w.Write(data)

}
