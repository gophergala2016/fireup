package main

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, pageStr)
}




func SlideHandler(w http.ResponseWriter, r *http.Request) {
}


func AssetHandler(w http.ResponseWriter, r *http.Request) {
}
