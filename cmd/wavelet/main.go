package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port = flag.String("port", "8080", "server port.")
	host := fmt.Sprintf(":%s", *port)

	http.HandleFunc("/", ok)
	http.HandleFunc("/points", points)

	log.Fatal(http.ListenAndServe(host, nil))
}

func ok(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "ok")
}
