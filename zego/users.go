package zego

func (a Auth) ListUsers() *Resource {

	path := "/api/v2/users.json"
	resource := api(a, "GET", path, "")

	return resource

}

func (a Auth) ShowUser(user_id string) *Resource {

	path := "/users/" + user_id + ".json"
	resource := api(a, "GET", path, "")

	return resource

}

func (a Auth) ShowUserRelated(user_id string) *Resource {

	path := "/users/" + user_id + "/related.json"
	resource := api(a, "GET", path, "")

	return resource

}
