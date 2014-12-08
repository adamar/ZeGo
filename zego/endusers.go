
package zego

type EndUser struct {
    Id              int `json:"id"`
    Url             string `json:"url"`
    Name            string `json:"name"`
    CreatedAt       string `json:"created_at"`
    UpdatedAt       string `json:"updated_at"`
    TimeZone        string `json:"time_zone"`
    Email           string `json:"email"`
    Phone           string `json:"phone"`
    Locale          string `json:"locale"`
    LocaleId        int `json:"locale_id"`
    OrganizationId  int `json:"organization_id"`
    Role            string `json:"role"`
    Verified        bool `json:"verified"`
    Photo           *Photo `json:"photo"`
}



type Photo struct {
    Id             int `json:"id"`
    Name           string `json:"name"`
    ContentUrl     string `json:"content_url"`
    ContentType    string `json:"content_type"`
    Size           int `json:"size"`
    Thumbnails     []*Thumbnail `json:"thumbnails"`
}



type Thumbnail struct {
    Id             int `json:"id"`
    Name           string `json:"name"`
    ContentUrl     string `json:"content_url"`
    ContentType    string `json:"content_type"`
    Size           int `json:"size"`
}
