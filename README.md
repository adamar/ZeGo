
ZeGo: Zendesk API wrapper for Golang
====================================

<p align="center">
  <img src="https://raw.githubusercontent.com/adamar/zego/master/doc/zendesk_logo.png" alt="Zendesk Logo"/>
</p>

## About:

A (hopefully) simple library for interfacing with the Zendesk v2 API

This library is in flux, so please use Godeps (or similar).

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

- GET `/users/{id}.json`
        
- GET `/users/{id}/related.json`


#### Tickets

- GET `/api/v2/tickets.json`

- DELETE `/api/v2/tickets/{id}.json`


#### Search

- GET `/api/v2/search.json?query={search_string}`


#### Organizations

- GET `/api/v2/organizations.json`

- GET `/api/v2/users/{user_id}/organizations.json`


#### Views

- GET `/api/v2/views.json`

- GET `/api/v2/active/views.json`

- GET `/api/v2/views/compact.json`

- GET `/api/v2/views/{id}.json


#### Macros

- GET `/api/v2/macros.json`


#### Triggers

- GET `/api/v2/triggers/active.json`


#### Satisfaction

- GET `/api/v2/satisfaction_ratings.json`

- GET `/api/v2/satisfaction_ratings/{id}.json`


#### Tags

- GET `/api/v2/tickets/{id}/tags.json`

- GET `/api/v2/topics/{id}/tags.json`

- GET `/api/v2/organizations/{id}/tags.json`


#### Groups

- GET `/api/v2/groups.json`

- GET `/api/v2/users/{user_id}/groups.json`

- GET `/api/v2/groups/assignable.json`

- GET `/api/v2/groups/{id}.json`


#### Incremental

- GET `/api/v2/incremental/tickets.json`


#### Comments

- GET `/api/v2/tickets/{ticket_id}/comments.json`
