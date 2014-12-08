package zego


type Macro struct {
    Id           int `json:"id"`
    Title        string `json:"title"`
    Active       bool `json:"active"`
    Actions      interface{} 
    Restriction  Macro_Restriction `json:"restriction"`
}

type Macro_Restriction struct {
    Type         string `json:"type"`
    Id           int `json:"id"`
}


func (a Auth) ListMacros() (*Resource, error) {

	path := "/macros.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) ApplyMacro(ticket_id string, macro_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + "/macros/" + macro_id + "/apply.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}
