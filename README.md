
ZeGo: Zendesk API wrapper for Golang
====================================

About:

A simple (hopefully) library for interfacing with the Zendesk v2 API


Usage:

```sh
auth := zego.Auth{"username", "password", "subdomain.zendesk.com"}
response, _ := auth.ListTickets()
tickets := &zego.TicketArray{}
json.Unmarshal([]byte(response.Raw), tickets)
```


