package zego

func (a Auth) ListTriggers() (*Resource, error) {

	path := "/triggers.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) ListActiveTriggers() (*Resource, error) {

	path := "/triggers/active.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}
