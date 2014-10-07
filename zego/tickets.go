package zego

func (a Auth) ListTickets() (*Resource, error) {

	path := "/tickets.json"
	resource, err := api(a, "GET", path, "")
        if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) GetTicket(ticket_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + ".json"
	resource, err := api(a, "GET", path, "")
        if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) GetMultipleTickets(ticket_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + ".json"
	resource, err := api(a, "GET", path, "")
        if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) GetTicketComments(ticket_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + "/comments.json"
	resource, err := api(a, "GET", path, "")
        if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) DeleteTicket(ticket_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + ".json"
	resource, err := api(a, "DELETE", path, "")
        if err != nil {
		return nil, err
	}

	return resource, nil

}
