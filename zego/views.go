package zego


View struct {
    Id           int `json:"id"`
    Title        string `json:"title"`
    Active       bool `json:"active"`
    SlaId        int `json:"sla_id"`
    Execution    []interface{} `json:"execution"`
    Conditions   []interface{} `json:"conditions"`
    Restriction  *Restriction `json:"restriction"`
}


Restriction  struct {
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
