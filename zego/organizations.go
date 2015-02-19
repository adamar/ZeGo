package zego

type OrganizationArray struct {
	Users []*Organization
}

type Organization struct {
	Id                 int                    `json:"id"`
	ExternalId         string                 `json:"external_id"`
	Url                string                 `json:"url"`
	Name               string                 `json:"name"`
	CreatedAt          string                 `json:"created_at"`
	UpdatedAt          string                 `json:"updated_at"`
	DomainNames        []string               `json:"domain_names"`
	Details            string                 `json:"details"`
	Notes              string                 `json:"notes"`
	GroupId            int                    `json:"group_id"`
	SharedTickets      bool                   `json:"shared_tickets"`
	SharedComments     bool                   `json:"shared_comments"`
	Tags               []string               `json:"tags"`
	OrganizationFields []*OrganizationalField `json:"organization_fields"`
}

type OrganizationField struct {
	OrgDropdown string  `json:"org_dropdown"`
	OrgDecimal  float32 `json:"org_decimal"`
}

func (a Auth) ListOrganizations() (*Resource, error) {

	path := "/organizations.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}
