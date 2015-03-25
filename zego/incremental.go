package zego

import (
	"encoding/json"
)

type IncrementalTicketArray struct {
	Count         int      `json:"count"`
	EndTime       uint     `json:"end_time"`
	Created       string   `json:"created"`
	Next_page     string   `json:"next_page"`
	Previous_page string   `json:"previous_page"`
	Tickets       []Ticket `json:"tickets"`
}

func (a Auth) IncrementalTicket(start_time string) (*IncrementalTicketArray, error) {

	TicketStruct := &IncrementalTicketArray{}

	path := "/incremental/tickets.json?start_time=" + start_time
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(resource.Raw), TicketStruct)

	return TicketStruct, nil

}
