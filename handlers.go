package main

import (
	"log"
	"fmt"
	"path"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	RootTemplate.Execute(w, nil)
}




func SlideHandler(w http.ResponseWriter, r *http.Request) {
	key := path.Base(r.URL.Path)

        log.Println(key)
        fmt.Fprintf(w, Get(key))

}


func AssetHandler(w http.ResponseWriter, r *http.Request) {
}
