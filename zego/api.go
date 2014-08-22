package zego

import (
	"log"
	//"net/url"
	"crypto/tls"
	"io/ioutil"
	"net/http"
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

func api(auth Auth, meth string, path string, params string) *Resource {

	trn := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: trn,
	}

	req, err := http.NewRequest(meth, "https://"+path, nil)
	errHandler(err)

	req.Header.Add("Content-Type", "application/json")

	req.SetBasicAuth(auth.Username, auth.Password)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	errHandler(err)

	data, err := ioutil.ReadAll(resp.Body)
	errHandler(err)

	return &Resource{Response: &resp, Raw: string(data)}

}
