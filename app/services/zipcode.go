package services

import (
	"io/ioutil"
	"net/http"
	"os"
)

type ZipData struct {
	Json   string
	Status int
}

func Zipcode(zipcode string) ZipData {

	jwt := os.Getenv("API_JWT_TOKEN")
	url := os.Getenv("API_URL") + "/shipments/zipcode/"
	url = url + zipcode

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("authorization", `Bearer `+jwt)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return ZipData{string(body), res.StatusCode}
}
