package zego

func (a Auth) ListViews() (*Resource, error) {

	path := "/views.json"
	resource, err := api(a, "GET", path, "")
        if err != nil {
		return nil, err
	}

	return resource, nil

}
