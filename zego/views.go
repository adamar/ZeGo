package zego


type View struct {
    Id           int `json:"id"`
    Title        string `json:"title"`
    Active       bool `json:"active"`
    SlaId        int `json:"sla_id"`
    Execution    []interface{} `json:"execution"`
    Conditions   []interface{} `json:"conditions"`
    Restriction  *Restriction `json:"restriction"`
}


type Restriction  struct {
    Type         string `json:"type"`
    Id           int `json:"id"`
} 



func (a Auth) ListViews() (*Resource, error) {

	path := "/views.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}



func (a Auth) ListActiveViews() (*Resource, error) {

    path := "/views/active.json"
    resource, err := api(a, "GET", path, "")
    if err != nil {
        return nil, err
    }

    return resource, nil

}


func (a Auth) ListViewsCompact() (*Resource, error) {

    path := "/views/compact.json"
    resource, err := api(a, "GET", path, "")
    if err != nil {
        return nil, err
    }

    return resource, nil

}



func (a Auth) GetView(ticket_id string) (*Resource, error) {

    path := "/views/" + ticket_id + ".json"
    resource, err := api(a, "GET", path, "")
    if err != nil {
        return nil, err
    }

    return resource, nil

}



