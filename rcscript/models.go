package rcscript

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CallEnterEvent struct {
	AccountId   string  `json:"accountId"`
	ExtensionId string  `json:"extensionId"`
	InParty     InParty `json:"inParty"`
	PartyId     string  `json:"partyId"`
	SessionId   string  `json:"sessionId"`
}

type InParty struct {
	Id   string    `json:"id"`
	From Extension `json:"from"`
	To   Extension `json:"to"`
}

type Extension struct {
	PhoneNumber string `json:"phoneNumber"`
}

type CallExitEvent struct {
	AccountId   string `json:"accountId"`
	ExtensionId string `json:"extensionId"`
	PartyId     string `json:"partyId"`
	SessionId   string `json:"sessionId"`
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
	AccountId   string `json:"accountId"`
	Command     string `json:"command"`
	CommandId   string `json:"commandId"`
	ExtensionId string `json:"extensionId"`
	PartyId     string `json:"partyId"`
	SessionId   string `json:"sessionId"`
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
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, evt)
}
