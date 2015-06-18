package zego

import (
	"encoding/json"
	"time"
)

type TicketField struct {
	Id                 uint64        `json:"id"`
	Url                string        `json:"url"`
	Type               string        `json:"type"`
	Title              string        `json:"title"`
	RawTitle           string        `json:"raw_title"`
	Description        string        `json:"description"`
	RawDescription     string        `json:"raw_description"`
	Position           int           `json:"position"`
	Active             bool          `json:"active"`
	Required           bool          `json:"required"`
	CollapsedForAgents bool          `json:"collapsed_for_agents"`
	RegexForValidation string        `json:"regexp_for_validation"`
	TitleInPortal      string        `json:"title_in_portal"`
	RawTitleInPortal   string        `json:"raw_title_in_portal"`
	VisibleInPortal    bool          `json:"visible_in_portal"`
	EditableInPortal   bool          `json:"editable_in_portal"`
	RequiredInPortal   bool          `json:"required_in_portal"`
	Tag                string        `json:"tag"`
	CreatedAt          time.Time     `json:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at"`
	SystemFieldOptions []interface{} `json:"system_field_options"`
	CustomFieldOptions []interface{} `json:"custom_field_options"`
	Removable          bool          `json:"removable"`
}

type TicketFieldResponse struct {
	TicketFields []TicketField `json:"ticket_fields"`
}

func (a Auth) ListFields() (*TicketFieldResponse, error) {
	ticketResp := &TicketFieldResponse{}

	path := "/ticket_fields.json"
	res, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(res.Raw), ticketResp)
	if err != nil {
		return nil, err
	}

	return ticketResp, nil
}
