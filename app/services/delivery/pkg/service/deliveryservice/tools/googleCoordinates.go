package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
)

type googleApiResponse struct {
	Results Results `json:"results"`
}

type Results []Geometry

type Geometry struct {
	Geometry Location `json:"geometry"`
}

type Location struct {
	Location Coordinates `json:"location"`
}

type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

func GetCoordinates(address string) (*Coordinates, error) {
	googleApiKey := configreader.GetString("googleApiKey")
	if googleApiKey == "" {
		return nil, errors.New("no google api uri in config")
	}

	resp, err := http.Get("https://maps.googleapis.com/maps/api/geocode/json?address=" + address + "&key=" + googleApiKey)
	if err != nil {
		return nil, fmt.Errorf("fetching google api uri data error: %q", err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	if err != nil {
		return nil, fmt.Errorf("reading google api data error: %q", err)
	}

	var data googleApiResponse
	if err = json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return &data.Results[0].Geometry.Location, nil
}
