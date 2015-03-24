package zego

import (
	"encoding/json"
)

type UserArray struct {
	Users []User
}

type SingleUser struct {
	User *User `json:"user"`
}

type User struct {
	Id                    uint32         `json:"id"`
	Url                   string         `json:"url"`
	Name                  string         `json:"name"`
	External_id           string         `json:"external_id"`
	Alias                 string         `json:"alias"`
	Created_at            string         `json:"created_at"`
	Updated_at            string         `json:"updated_at"`
	Active                bool           `json:"active"`
	Verified              bool           `json:"verified"`
	Shared                bool           `json:"shared"`
	Shared_agent          bool           `json:"shared_agent"`
	Locale                string         `json:"locale"`
	Locale_id             uint32         `json:"locale_id"`
	Time_zone             string         `json:"time_zone"`
	Last_login_at         string         `json:"last_login_at"`
	Email                 string         `json:"email"`
	Phone                 string         `json:"phone"`
	Signature             string         `json:"signature"`
	Details               string         `json:"details"`
	Notes                 string         `json:"notes"`
	Organization_id       uint32         `json:"organization_id"`
	Role                  string         `json:"role"`
	Customer_role_id      uint32         `json:"custom_role_id"`
	Moderator             bool           `json:"moderator"`
	Ticket_restriction    string         `json:"ticket_restriction"`
	Only_private_comments bool           `json:"only_private_comments"`
	Tags                  []string       `json:"tags"`
	Restricted_agent      bool           `json:"restricted_agent"`
	Suspended             bool           `json:"suspended"`
	Photo                 []*Users_Photo `json:"photo"`
	User_fields           []*User_Field  `json:"user_fields"`
}

type Users_Photo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ContentUrl  string `json:"content_url"`
	ContentType string `json:"content_type"`
	Size        int    `json:"size"`
}

type User_Field struct {
	UserDecimal  float32 `json:"user_decimal"`
	UserDropdown string  `json:"user_dropdown"`
	UserDate     string  `json:"user_date"`
}

func (a Auth) ListUsers() (*Resource, error) {

	path := "/users.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) ShowUser(user_id string) (*SingleUser, error) {

	UserStruct := &SingleUser{}

	path := "/users/" + user_id + ".json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(resource.Raw), UserStruct)

	return UserStruct, nil

}

func (a Auth) ShowUserRelated(user_id string) (*Resource, error) {

	path := "/users/" + user_id + "/related.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) ListCollaborators(ticket_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + "/collaborators.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}
