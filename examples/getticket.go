package main

import (
	"github.com/adamar/zego"
	"log"
)

func main() {

	auth := zego.Auth{"user@example.com", "password", "subdomain.zendesk.com"}
        tickets := auth.ListTickets()
	log.Print(tickets)

}
