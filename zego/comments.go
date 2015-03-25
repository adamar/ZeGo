package zego

import (
	"encoding/json"
)

type CommentArray struct {
	Comments []Comments `json:"comments"`
}

type Comments struct {
	Id          int           `json:"id"`
	Type        string        `json:"type"`
	Body        string        `json:"body"`
	Public      bool          `json:"public"`
	CreatedAt   string        `json:"created_at"`
	AuthorId    int           `json:"author_id"`
	Attachments []Attachments `json:"attachments"`
	Metadata    interface{}   `json:"metadata"`
	Via         interface{}   `json:"via"`
}

type Attachments struct {
	Id          int           `json:"id"`
	Name        string        `json:"name"`
	ContentURL  string        `json:"content_url"`
	ContentType string        `json:"content_type"`
	Size        int           `json:"size"`
	Thumbnails  []interface{} `json:"thumbnails"`
}

func (a Auth) ListComments(ticket_id string) (*CommentArray, error) {

	CommentStruct := &CommentArray{}

	path := "/tickets/" + ticket_id + "/comments.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(resource.Raw), CommentStruct)

	return CommentStruct, nil

}
