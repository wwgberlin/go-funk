# wavelet

Run the server with:

    make run-server

You can get the waveform points with:

    curl http://localhost:8080/points

Try rendering a couple waveforms with:

    curl http://localhost:8080/waveform > waveform.png
    curl http://localhost:8080/waveform?width=1200&colors=reds&sampling=abs_avg > waveform.png

## workshop

1. Reduce all samples to a smaller size. (from 18M to 1200)
1. Render samples as PNG waveform.
1. Render using different colors depending on the value.

## Other ideas

* Run FFT on samples to have an histogram.
* Render a GIF of the histogram over time. (synchronized with the track)
* Render a GIF of the samples over time. (synchronized with the track)
* Handle POST request of WAV data and render them "in real time".
