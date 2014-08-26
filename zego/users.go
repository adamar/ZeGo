package zego

func ListUsers(auth Auth) *Resource {

	path := auth.Subdomain + "/api/v2/users.json"
	resource := api(auth, "GET", path, "")

	return resource

}

func ShowUser(auth Auth, user_id string) *Resource {

	path := auth.Subdomain + "/api/v2/users/" + user_id + ".json"
	resource := api(auth, "GET", path, "")

	return resource

}

func ShowUserRelated(auth Auth, user_id string) *Resource {

	path := auth.Subdomain + "/api/v2/users/" + user_id + "/related.json"
	resource := api(auth, "GET", path, "")

	return resource

}
