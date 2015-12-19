package zego

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTicket(t *testing.T) {

	JSON := `
    {"tickets":[{
     "id":               35436,
     "url":              "https://company.zendesk.com/api/v2/tickets/35436.json",
     "external_id":      "ahg35h3jh",
     "created_at":       "2009-07-20T22:55:29Z",
  "updated_at":       "2011-05-05T10:38:52Z",
  "type":             "incident",
  "subject":          "Help, my printer is on fire!",
  "raw_subject":      "{{dc.printer_on_fire}}",
  "description":      "The fire is very colorful.",
  "priority":         "high",
  "status":           "open",
  "recipient":        "support@company.com",
  "requester_id":     20978392,
  "submitter_id":     76872,
  "assignee_id":      235323,
  "organization_id":  509974,
  "group_id":         98738,
  "collaborator_ids": [35334, 234],
  "forum_topic_id":   72648221,
  "problem_id":       9873764,
  "has_incidents":    false,
  "due_at":           null,
  "tags":             ["enterprise", "other_tag"],
  "via": {
    "channel": "web"
  },
  "custom_fields": [
    {
      "id":    27642,
      "value": "745"
    },
    {
      "id":    27648,
      "value": "yes"
    }
  ],
  "satisfaction_rating": {
    "id": 1234,
    "score": "good",
    "comment": "Great support!"
  },
  "sharing_agreement_ids": [84432]
},
{
  "id":               35436,
  "url":              "https://company.zendesk.com/api/v2/tickets/35436.json",
  "external_id":      "ahg35h3jh",
  "created_at":       "2009-07-20T22:55:29Z",
  "updated_at":       "2011-05-05T10:38:52Z",
  "type":             "incident",
  "subject":          "Help, my printer is on fire!",
  "raw_subject":      "{{dc.printer_on_fire}}",
  "description":      "The fire is very colorful.",
  "priority":         "high",
  "status":           "open",
  "recipient":        "support@company.com",
  "requester_id":     20978392,
  "submitter_id":     76872,
  "assignee_id":      235323,
  "organization_id":  509974,
  "group_id":         98738,
  "collaborator_ids": [35334, 234],
  "forum_topic_id":   72648221,
  "problem_id":       9873764,
  "has_incidents":    false,
  "due_at":           null,
  "tags":             ["enterprise", "other_tag"],
  "via": {
    "channel": "web"
  },
  "custom_fields": [
    {
      "id":    27642,
      "value": "745"
    },
    {
      "id":    27648,
      "value": "yes"
    }
  ],
  "satisfaction_rating": {
    "id": 1234,
    "score": "good",
    "comment": "Great support!"
  },
  "sharing_agreement_ids": [84432]
}],"next_page":"https://company.zendesk.com/api/v2/tickets.json?page=2","previous_page":null,"count":2}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, JSON)
	}))
	defer ts.Close()

	auth := Auth{"user@example.com", "Password", ts.URL}
	response, err := auth.ListTickets()
	if err != nil {
		t.Error(err)
	}

	if response.Tickets[0].Id != 35436 {
		t.Error("Ticket JSON unmarshalling error")
	}

	wantFld := []struct {
		id    uint64
		value string
	}{
		{27642, "745"},
		{27648, "yes"},
	}
	for i, fld := range response.Tickets[0].Custom_Fields {
		if fld.Id != wantFld[i].id {
			t.Errorf("Wanted custom field id %d, have %d", wantFld[i].id, fld.Id)
		}
		val, ok := fld.Value.(string)
		if !ok {
			t.Errorf("Value of custom field %d is not string but %T", i, fld.Value)
		}
		if val != wantFld[i].value {
			t.Errorf("Wanted custom field value %s, have %s", wantFld[i].value, fld.Value)
		}
	}

}
