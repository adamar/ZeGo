package zego

func ListViews(auth Auth) *Resource {

	path := auth.Subdomain + "/api/v2/" + "/views.json"
	resource := api(auth, "GET", path, "")

	return resource

}
