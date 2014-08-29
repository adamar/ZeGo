package zego

func (a Auth) ListViews() *Resource {

	path := "/views.json"
	resource := api(a, "GET", path, "")

	return resource

}
