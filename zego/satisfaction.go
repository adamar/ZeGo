
package zego

type Satisfaction_Ratings struct {
    Ratings            []*Satisfaction
}

type Satisfaction struct {
    Id                 int `json:"id"`
    Url                string `json:"url"`
    AssigneeId         int `json:"assignee_id"`
    GroupId            int `json:"group_id"`
    RequesterId        int `json:"requester_id"`
    TicketId           int `json:"ticket_id"`
    Score              string `json:"score"`
    UpdatedAt          string `json:"updated_at"`
    CreatedAt          string `json:"created_at"`
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



