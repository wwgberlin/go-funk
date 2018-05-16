# wavelet

In this challenge we will render images and gifs of waveforms from a wav file.

The following functionality has been removed and needs to be implmeneted:

- `AbsAvg` - A sampler function that samples a slice and returns the average values of the slice's absolute values
- `Project` - function that normalizes the values in a slice of integers to a given maximum
- `DrawRectangle` - A function that draws a rectangle of given bounds inside an image.
- `RenderGopherFrame` A function that returns a paletted image for a gif
- `ColorGopherFunc` - A closure ColorFunc function that returns the color of a pixel in the original image. Given new bounds.

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
