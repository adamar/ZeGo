package zego

func (a Auth) ListViews() *Resource {

	path := a.Subdomain + "/api/v2/" + "/views.json"
	resource := api(a, "GET", path, "")

	return resource

}
