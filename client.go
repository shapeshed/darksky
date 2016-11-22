package darksky

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

const BASEURL = "https://api.darksky.net/forecast/"

var debug = os.Getenv("DEBUG")

type RequestParams struct {
	Key       string
	Latitude  float64
	Longitude float64
	Exclude   string
	Extend    string
	Lang      string
	Units     string
}

func latLongToString(lat, long float64) string {
	longString := strconv.FormatFloat(lat, 'f', 6, 64)
	latString := strconv.FormatFloat(long, 'f', 6, 64)
	return longString + "," + latString
}

//TODO - return missing params error
// Required are r.Key, r.Latitude, r.Longitude
func makeURL(r *RequestParams) (string, error) {

	u, err := url.Parse(BASEURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path,
		r.Key,
		latLongToString(r.Latitude, r.Longitude),
	)

	params := url.Values{}

	if r.Exclude != "" {
		params.Add("exclude", r.Exclude)
	}
	if r.Extend != "" {
		params.Add("extend", r.Extend)
	}
	if r.Lang != "" {
		params.Add("lang", r.Lang)
	}
	if r.Units != "" {
		params.Add("units", r.Units)
	}

	u.RawQuery = params.Encode()

	return u.String(), nil

}

func Get(p *RequestParams) (forecast *Forecast, err error) {

	url, err := makeURL(p)
	if err != nil {
		return forecast, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return forecast, err
	}
	req.Header.Add("Accept-Encoding", "gzip")

	res, err := client.Do(req)
	if err != nil {
		return forecast, err
	}
	if res != nil {
		defer res.Body.Close()
	}

	if res.StatusCode != 200 {
		return forecast, fmt.Errorf("Received %v response", res.StatusCode)
	}

	body, err := gzip.NewReader(res.Body)
	if err != nil {
		return forecast, err
	}
	defer body.Close()

	decoder := json.NewDecoder(body)

	err = decoder.Decode(&forecast)
	if err != nil {
		return forecast, err
	}

	return forecast, err

}
