package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	var port = flag.String("port", "8080", "server port.")
	addr := net.JoinHostPort("", *port)

	http.HandleFunc("/", ok)
	http.HandleFunc("/points", points)
	http.HandleFunc("/waveform", waveform)
	http.HandleFunc("/gif", gifHandler)
	http.HandleFunc("/gif2", gifHandler2)

	log.Fatal(http.ListenAndServe(addr, nil))
}

func ok(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "ok")
}
