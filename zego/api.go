package zego

import (
	"io/ioutil"
	"log"
	"net/http"
	//"time"
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

	req, err := http.NewRequest(meth, "https://"+auth.Subdomain+"/api/v2/"+path, nil)
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
