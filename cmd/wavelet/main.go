package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/soundcloud/wavelet/lib/sampler"
	"github.com/soundcloud/wavelet/lib/wav"
)

func main() {
	var port = flag.String("port", "8080", "server port.")
	host := fmt.Sprintf(":%s", *port)

	http.HandleFunc("/", ok)

	log.Fatal(http.ListenAndServe(host, nil))
}

func ok(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("fixtures/fixture.wav")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	samples, err := wav.Samples(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	points, err := sampler.Sample(samples, 1200)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, fmt.Sprintf("%v", points))
}
