package zego

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Resource struct {
	//Headers     http.Header
	Response interface{}
	Raw      string
}

type Auth struct {
	Username  string
	Password  string
	Subdomain string
}

func errHandler(err error) {
	if err != nil {
		log.Print(err)
	}
}

func api(auth Auth, meth string, path string, params string) (*Resource, error) {

	trn := &http.Transport{}

	client := &http.Client{
		Transport: trn,
	}

	var URL string

	// Check if entire URL is in path
	if strings.HasPrefix(path, "http") {
		URL = path

		// Otherwise build url from auth components
	} else {
		if strings.HasPrefix(auth.Subdomain, "http") {
			URL = auth.Subdomain + "/api/v2/" + path
		} else {
			URL = "https://" + auth.Subdomain + "/api/v2/" + path
		}
	}

	req, err := http.NewRequest(meth, URL, bytes.NewBufferString(params))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	req.SetBasicAuth(auth.Username, auth.Password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Resource{Response: &resp, Raw: string(data)}, nil

}
