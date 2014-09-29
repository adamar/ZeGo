package zego

func (a Auth) ListMacros() *Resource {

	path := "/macros.json"
	resource := api(auth, "GET", path, "")

	return resource

}

func (a Auth) ApplyMacro(ticket_id string, macro_id string) *Resource {

	path := "/tickets/" + ticket_id + "/macros/" + macro_id + "/apply.json"
	resource := api(auth, "GET", path, "")

	return resource

}
