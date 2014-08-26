package zego

func ListMacros(auth Auth) *Resource {

	path := auth.Subdomain + "/api/v2/macros.json"
	resource := api(auth, "GET", path, "")

	return resource

}



func ApplyMacro(auth Auth, ticket_id string, macro_id string) *Resource {

	path := auth.Subdomain + "/api/v2/" + "/tickets/" + ticket_id + "/macros/" + macro_id + "/apply.json"
	resource := api(auth, "GET", path, "")

	return resource

}
