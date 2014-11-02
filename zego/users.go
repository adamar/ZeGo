package zego

type User struct {
        Id                    uint32
        Url                   string
        Name                  string
        External_id           string
        Alias                 string
        Created_at            string
        Updated_at            string
        Active                bool
        Verified              bool
        Shared                bool
        Shared_agent          bool
        Locale                string
        Locale_id             uint32
        Time_zone             
        Last_login_at
        Email
        Phone
        Signature
        Details
        Notes
        Organization_id
        Role
        Customer_role_id
        Moderator
        Ticket_restriction
        Only_private_comments
        Tags
        Restricted_agent
        Suspended
        Photo
        User_fields
}



func (a Auth) ListUsers() (*Resource, error) {

	path := "/users.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) ShowUser(user_id string) (*Resource, error) {

	path := "/users/" + user_id + ".json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) ShowUserRelated(user_id string) (*Resource, error) {

	path := "/users/" + user_id + "/related.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}
