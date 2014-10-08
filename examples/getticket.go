package main

import (
	"github.com/adamar/zego"
	"log"
)

func main() {

	auth := zego.Auth{"user@example.com", "password", "subdomain.zendesk.com"}
        response := auth.ListTickets()
        tickets := &zego.TicketArray{}
        json.Unmarshal([]byte(response.Raw), tickets)
	log.Print(tickets)

}
