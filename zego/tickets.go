package zego

func (a Auth) ListTickets() *Resource {

	path := "/tickets.json"
	resource := api(a, "GET", path, "")

	return resource

}

func GetTicket(auth Auth, ticket_id string) *Resource {

	path := "/api/v2/" + "/tickets/" + ticket_id + ".json"
	resource := api(auth, "GET", path, "")

	return resource

}

func GetMultipleTickets(auth Auth, ticket_id string) *Resource {

	path := "/api/v2/" + "/tickets/" + ticket_id + ".json"
	resource := api(auth, "GET", path, "")

	return resource

}

func GetTicketComments(auth Auth, ticket_id string) *Resource {

	path := "/api/v2/" + "/tickets/" + ticket_id + "/comments.json"
	resource := api(auth, "GET", path, "")

	return resource

}

func DeleteTicket(auth Auth, ticket_id string) *Resource {

	path := "/api/v2/" + "/tickets/" + ticket_id + ".json"
	resource := api(auth, "DELETE", path, "")

	return resource

}
