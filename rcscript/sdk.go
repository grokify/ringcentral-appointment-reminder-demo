package rcscript

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/net/httputilmore"
	"github.com/grokify/mogo/net/urlutil"
)

const (
	CommandPlay     string = "Play"
	StatusCompleted string = "Completed"
)
const (
	DefaultParamValue                     string = "~"
	URLTelephonySessionFormat             string = "/restapi/v1.0/account/%s/telephony/sessions/%s"
	URLTelephonySessionsPartiesPlayFormat string = `/restapi/v1.0/account/%s/telephony/sessions/%s/parties/%s/play`
)

//https://platform.devtest.ringcentral.com/restapi/v1.0/account/accountId/telephony/sessions/telephonySessionId

// RcScriptSdk is a simple SDK for making Call Scripting Commands
type RcScriptSdk struct {
	ServerURL string
	Token     string
}

// Play plays a media file
func (sdk *RcScriptSdk) Play(sessionID, partyID string, body PlayRequest) (*http.Response, error) {
	apiUrl := urlutil.JoinAbsolute(sdk.ServerURL,
		fmt.Sprintf(URLTelephonySessionsPartiesPlayFormat, DefaultParamValue, sessionID, partyID))
	fmt.Println(http.MethodPost + " " + apiUrl)
	fmtutil.PrintJSON(body)

	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Add(httputilmore.HeaderAuthorization, "Bearer "+sdk.Token)
	req.Header.Add(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJSONUtf8)
	return client.Do(req)
}

// Hangsup a call
func (sdk *RcScriptSdk) Hangup(sessionID string) (*http.Response, error) {
	apiUrl := urlutil.JoinAbsolute(sdk.ServerURL,
		fmt.Sprintf(URLTelephonySessionFormat, DefaultParamValue, sessionID))
	fmt.Println(apiUrl)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, apiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(httputilmore.HeaderAuthorization, "Bearer "+sdk.Token)
	return client.Do(req)
}
