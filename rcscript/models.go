package rcscript

import (
	"encoding/json"
	"io"
	"net/http"
)

type CallEnterEvent struct {
	AccountID   string  `json:"accountId"`
	ExtensionID string  `json:"extensionId"`
	InParty     InParty `json:"inParty"`
	PartyID     string  `json:"partyId"`
	SessionID   string  `json:"sessionId"`
}

type InParty struct {
	ID   string    `json:"id"`
	From Extension `json:"from"`
	To   Extension `json:"to"`
}

type Extension struct {
	PhoneNumber string `json:"phoneNumber"`
}

type CallExitEvent struct {
	AccountID   string `json:"accountId"`
	ExtensionID string `json:"extensionId"`
	PartyID     string `json:"partyId"`
	SessionID   string `json:"sessionId"`
}

/*
	{
	  "accountId": "400131801008",
	  "extensionId": "400137552008",
	  "inParty": {
	    "from": {
	      "phoneNumber": "+12127150355"
	    },
	    "id": "p-467aa7237b524387b053c5d5f06787f5",
	    "to": {
	      "phoneNumber": "+12014320001"
	    }
	  },
	  "partyId": "p-ec4202d7dd35401f9e6230d428493701",
	  "sessionId": "s-82eecd460b564de1a42b395eb845d912"
	}
*/
type CommandUpdateEvent struct {
	AccountID   string `json:"accountId"`
	Command     string `json:"command"`
	CommandID   string `json:"commandId"`
	ExtensionID string `json:"extensionId"`
	PartyID     string `json:"partyId"`
	SessionID   string `json:"sessionId"`
	Status      string `json:"status"`
}

/*
{
	"accountId" : "400131801008",
	"command" : "Play",
	"commandId" : "171888866234845407",
	"extensionId" : "400137552008",
	"partyId" : "p-467aa7237b524387b053c5d5f06787f5",
	"sessionId" : "s-82eecd460b564de1a42b395eb845d912",
 "status" : "Completed"
 }
*/

func Bind(evt interface{}, r *http.Request) error {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, evt)
}
