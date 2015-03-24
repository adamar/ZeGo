package zego

import (
	"encoding/json"
)

func (a Auth) IncrementalTicket(start_time string) (*TicketArray, error) {

	TicketStruct := &TicketArray{}

	path := "/incremental/tickets.json?start_time=" + start_time
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(resource.Raw), TicketStruct)

	return TicketStruct, nil

}
