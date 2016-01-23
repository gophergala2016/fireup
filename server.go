package main

import (
	"net/http"
	"log"
	"fmt"
)

const pageStr = `
<html>
<head></head>
<body>
	<p>It works ...</p>
</body>
</html>
`

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, pageStr)
}

func StartHttpListener(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("starting http listener on http://%s\n", addr)

	http.HandleFunc("/", RootHandler)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Println("ERROR:", err)
	}
}
