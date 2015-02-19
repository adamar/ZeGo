package zego

type TagArray struct {
	Tags []string `json:"tags"`
}

func (a Auth) ShowTicketTags(ticket_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + "/tags.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) ShowTopicTags(topic_id string) (*Resource, error) {

	path := "/topics/" + topic_id + "/tags.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) ShowOrganizationTags(organization_id string) (*Resource, error) {

	path := "/organizations/" + organization_id + "/tags.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}
