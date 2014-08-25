package zego

func ListingTickets(auth Auth) *Resource {

	path := auth.Subdomain + "/api/v2/tickets.json"
	resource := api(auth, "GET", path, "")

	return resource

}

func GetTicket(auth Auth, ticket_id string) *Resource {

	path := auth.Subdomain + "/api/v2/" + "/tickets/" + ticket_id + ".json"
	resource := api(auth, "GET", path, "")

	return resource

}

func GetMultipleTickets(auth Auth, ticket_id string) *Resource {

	path := auth.Subdomain + "/api/v2/" + "/tickets/" + ticket_id + ".json"
	resource := api(auth, "GET", path, "")

	return resource

}
