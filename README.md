
ZeGo: Zendesk API wrapper for Golang
====================================

<p align="center">
  <img src="https://raw.githubusercontent.com/adamar/zego/master/doc/zendesk_logo.png" alt="Zendesk Logo"/>
</p>


## About:

A simple (hopefully) library for interfacing with the Zendesk v2 API


## Usage:

```sh
auth := zego.Auth{"username", "password", "subdomain.zendesk.com"}
response, _ := auth.ListTickets()
tickets := &zego.TicketArray{}
json.Unmarshal([]byte(response.Raw), tickets)
```


## Endpoints Implemented 


#### Users

- GET `/api/v2/users.json`


#### Tickets

- GET `/api/v2/tickets.json`

- DELETE `/api/v2/tickets/{id}.json`


#### Views

- GET `/api/v2/views.json`


