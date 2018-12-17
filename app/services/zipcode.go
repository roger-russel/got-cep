package services

import (
	"fmt"
	"os"
	"strconv"

	"github.com/imroc/req"
)

type sZipcode struct {
	name         string
	city         string
	neighborhood string
	state_short  string
	street       string
	zipcode      string
}

func Zipcode(zipcode int) string {

	jwt := os.Getenv("API_JWT_TOKEN")
	url := os.Getenv("API_URL") + "/shipments/zipcode/"
	url = url + strconv.Itoa(zipcode)

	req.Debug = true

	r, _ := req.Get(url, req.Header{
		"Accept":        "application/json",
		"Authorization": `Bearer ` + jwt,
	})

	resp := r.Response()
	fmt.Println(resp.StatusCode)
	return ""
	//return json.NewDecoder(resp.Body).Decode(target)
}
