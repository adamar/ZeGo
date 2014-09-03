package zego

import (
	"log"
	"crypto/tls"
	"io/ioutil"
	"net/http"
)


type Ticket struct {
  Id                 uint32
  Url                string
  External_id        string
  Created_at         time.Time
  Updated_at         time.Time
  Ttype              string
  Subject            string
  Raw_subject        string
  Description        string
  Priority           string
  Status             string
  Recipient          string
  Requester_id       uint32
  Submitter_id       uint32
  Assignee_id        uint32
  Organization_id    uint32
  Group_id           uint32
  Collaborator_ids   []uint32
  Forum_topic_id     uint32
  Problem_id         uint32
  Has_incidents      bool
  Due_at             time.Time
  Tags               []string
  Via                interface{}
}


type Connector interface {
    ListViews() *Resource
}


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

	req, err := http.NewRequest(meth, "https://" + auth.Subdomain + "/api/v2/" + path, nil)
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
