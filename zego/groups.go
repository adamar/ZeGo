package zego

type Group struct {
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Id        int    `json:"id"`
}

type Groups struct {
	Groups []*Group `json:"groups"`
}

func (a Auth) GetGroups() (*Resource, error) {

	path := "/groups.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) GetUserGroups(user_id string) (*Resource, error) {

	path := "/users/" + user_id + "/groups.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) GetAssignableGroups() (*Resource, error) {

	path := "/groups/assignable.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) GetGroup(group_id string) (*Resource, error) {

	path := "/groups/" + group_id + ".json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}
