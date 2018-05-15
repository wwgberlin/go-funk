# wavelet

Run the server with:

    make run-server

Or: 

    go test ./... && go build && ./go-funk
You can see all the samples at:

    curl http://localhost:8080/samples

Try rendering a couple waveforms with:

    http://localhost:8080/waveform
    http://localhost:8080/waveform?colors=reds

Try rendering a "real-time" gif of the samples:

    http://localhost:8080/gif?count=5000&width=50&height=200
    http://localhost:8080/gif2?count=5000&width=50&height=200
    http://localhost:8080/gopher
    
To see the *almost* synched version of the gif with the song:

    http://127.0.0.1:8080/rick/rick1.html
    http://127.0.0.1:8080/rick/rick2.html
