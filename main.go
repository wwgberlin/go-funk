package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/wwgberlin/go-funk/audio"
	"github.com/wwgberlin/go-funk/renderer"
	"github.com/wwgberlin/go-funk/sampler"
)

const filePath = "public/rick/fixture.wav"

func main() {
	var port = flag.String("port", "8080", "server port")
	flag.Parse()

	addr := net.JoinHostPort("", *port)

	sampler.RegSampler("", sampler.AbsAvg)
	renderer.RegColorFunc("", renderer.Colors["black"])

	data, err := getFileData(filePath)
	if err != nil {
		panic(err)
	}

	start(addr, data)
}

func start(addr string, data audio.WavData) {

	http.HandleFunc("/", status)
	http.HandleFunc("/samples", samplesHandler(data))
	http.HandleFunc("/waveform", waveformHandler(data))
	http.HandleFunc("/gif", gifHandler(data))
	http.HandleFunc("/gif2", gif2Handler(data))

	http.HandleFunc("/gopher", gopherHandler(data, gopherImage()))

	http.Handle("/rick/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))

	log.Fatal(http.ListenAndServe(addr, nil))
}

func getFileData(filePath string) (data audio.WavData, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	data, err = audio.NewWavData(file)
	return
}

func gopherImage() image.Image {
	f, err := os.Open("public/rick/gopher.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}
	return img

}
func status(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "ok")
}
