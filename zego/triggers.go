package zego

func ListTriggers(auth Auth) *Resource {

	path := auth.Subdomain + "/api/v2/" + "/triggers.json"
	resource := api(auth, "GET", path, "")

	return resource

}

func ListActiveTriggers(auth Auth) *Resource {

	path := auth.Subdomain + "/api/v2/" + "/triggers/active.json"
	resource := api(auth, "GET", path, "")

	return resource

}
