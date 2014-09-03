package zego

func (a Auth) ListTickets() *Resource {

	path := "/tickets.json"
	resource := api(a, "GET", path, "")

	return resource

}

func (a Auth) GetTicket(ticket_id string) *Resource {

	path := "/tickets/" + ticket_id + ".json"
	resource := api(a, "GET", path, "")

	return resource

}

func (a Auth) GetMultipleTickets(ticket_id string) *Resource {

	path := "/tickets/" + ticket_id + ".json"
	resource := api(a, "GET", path, "")

	return resource

}

func (a Auth) GetTicketComments(ticket_id string) *Resource {

	path := "/tickets/" + ticket_id + "/comments.json"
	resource := api(a, "GET", path, "")

	return resource

}

func (a Auth) DeleteTicket(ticket_id string) *Resource {

	path := "/tickets/" + ticket_id + ".json"
	resource := api(a, "DELETE", path, "")

	return resource

}
