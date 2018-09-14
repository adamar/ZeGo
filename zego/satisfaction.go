package zego

type Satisfaction_Ratings struct {
	Ratings []*Satisfaction
}

type Satisfaction struct {
	Id          int64  `json:"id"`
	Url         string `json:"url"`
	AssigneeId  int64  `json:"assignee_id"`
	GroupId     int64  `json:"group_id"`
	RequesterId int64  `json:"requester_id"`
	TicketId    int64  `json:"ticket_id"`
	Score       string `json:"score"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
}

func (a Auth) ListSatisfactionRatings() (*Resource, error) {

	path := "/satisfaction_ratings.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) GetSatisfactionRating(rating_id string) (*Resource, error) {

	path := "/satisfaction_ratings/" + rating_id + ".json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}
