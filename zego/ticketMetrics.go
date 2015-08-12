package zego

//GET /api/v2/tickets/{ticket_id}/metrics.json

type TicketMetric struct {
    Id                 int `json:"id"`
    Ticket_id          int `json:"ticket_id"`
    Replies            int `json:"replies"`
    ReplyTime          int `json:"reply_time_in_minutes.business"`
    FirstResolution	   int `json:"first_resolution_time_in_minutes.business"`
    CreatedAt          string `json:"created_at"`
}


func (a Auth) GetTicketMetrics(ticket_id int) (*Resource, error) {

    path := "/tickets/" + string(ticket_id) + "/metrics.json"
    resource, err := api(a, "GET", path, "")
    if err != nil {
        return nil, err
    }
    return resource, nil
}