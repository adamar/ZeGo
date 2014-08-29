package zego

func ListTriggers(auth Auth) *Resource {

	path := "/api/v2/" + "/triggers.json"
	resource := api(auth, "GET", path, "")

	return resource

}

func ListActiveTriggers(auth Auth) *Resource {

	path := "/api/v2/" + "/triggers/active.json"
	resource := api(auth, "GET", path, "")

	return resource

}
