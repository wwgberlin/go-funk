package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/wwgberlin/go-funk/renderer"
	"github.com/wwgberlin/go-funk/sampler"
)

type RequestData struct {
	Conf    renderer.Config
	Sampler sampler.SamplerFunc
}

func parseInt(str string, def int, prevErr error) (i int, err error) {
	if prevErr != nil {
		return i, prevErr
	}
	if str == "" {
		return def, nil
	}
	return strconv.Atoi(str)
}

func parseRequest(req *http.Request, defaults *RequestData) (*RequestData, error) {
	var (
		d   RequestData
		err error
	)

	query := req.URL.Query()

	if colorKey := query.Get("colors"); colorKey == "" {
		d.Conf.Colorer = defaults.Conf.Colorer
	} else {
		if colorFunc, ok := renderer.Colors[colorKey]; !ok {
			return nil, fmt.Errorf("could not find color method with key %s", colorKey)
		} else {
			d.Conf.Colorer = colorFunc
		}
	}

	if samplingKey := req.URL.Query().Get("sampling"); samplingKey == "" {
		d.Sampler = defaults.Sampler
	} else {
		samplingFunc, ok := sampler.Samplers[samplingKey]
		if !ok {
			return nil, fmt.Errorf("could not find sampling method with key %s", samplingKey)
		} else {
			d.Sampler = samplingFunc
		}
	}

	d.Conf.Width, err = parseInt(query.Get("width"), defaults.Conf.Width, nil)
	d.Conf.Height, err = parseInt(query.Get("height"), defaults.Conf.Height, err)
	d.Conf.Count, err = parseInt(query.Get("count"), defaults.Conf.Count, err)

	return &d, err
}
