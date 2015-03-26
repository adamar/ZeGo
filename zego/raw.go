package zego

func (a Auth) RawGet(path string) (string, error) {

	resource, err := api(a, "GET", path, "")
	if err != nil {
		return "", err
	}

	return resource.Raw, err

}
