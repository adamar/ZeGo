package main

import (
	"encoding/json"
	"github.com/adamar/zego/zego"
	"log"
)

func main() {

	auth := zego.Auth{"user@example.com", "Password", "example.zendesk.com"}
	// auth := zego.Auth{"user@example.com/token", "Token", "example.zendesk.com"}

	response, _ = auth.ListOrganizations()
	orgs := &zego.OrganizationArray{}
	json.Unmarshal([]byte(response.Raw), orgs)
	for _, o := range orgs.Organizations {
		log.Print(o.Name)
	}

}
