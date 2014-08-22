package main

import (
	"github.com/adamar/zego"
	"log"
)

func main() {

	auth := zego.Auth{"user@example.com", "password", "subdomain.zendesk.com"}
        tickets := zego.GetTicket(auth,  "4542")
	log.Print(tickets)

}
