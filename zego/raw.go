package zego

import (
	"encoding/json"
)

func (a Auth) rawGet(path string) (string, error) {

	resource, err := api(a, "GET", path, "")
	if err != nil {
		return "", err
	}

	return resource.Raw, err

}
