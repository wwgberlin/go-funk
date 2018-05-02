package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	var port = flag.String("port", "8080", "server port.")
	host := fmt.Sprintf(":%s", *port)

	http.HandleFunc("/", ok)

	log.Fatal(http.ListenAndServe(host, nil))
}

func ok(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "ok!")
}
