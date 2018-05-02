package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

func main() {
	var port = flag.String("port", "8080", "server port.")

	http.HandleFunc("/", ok)

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func ok(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "ok!")
}
