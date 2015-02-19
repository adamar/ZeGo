package main

import (
	"encoding/json"
	"github.com/adamar/zego/zego"
	"log"
)

func main() {

	auth := zego.Auth{"user@example.com", "Password", "example.zendesk.com"}
	// auth := zego.Auth{"user@example.com/token", "Token", "example.zendesk.com"}

	response, _ := auth.ListTickets()
	tickets := &zego.TicketArray{}
	json.Unmarshal([]byte(response.Raw), tickets)
	for _, t := range tickets.Tickets {
		log.Print(t.Subject)
	}

}
