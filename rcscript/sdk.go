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
	UrlTelephonySessionFormat             string = "/restapi/v1.0/account/%s/telephony/sessions/%s"
	UrlTelephonySessionsPartiesPlayFormat string = `/restapi/v1.0/account/%s/telephony/sessions/%s/parties/%s/play`
)

//https://platform.devtest.ringcentral.com/restapi/v1.0/account/accountId/telephony/sessions/telephonySessionId

// RcScriptSdk is a simple SDK for making Call Scripting Commands
type RcScriptSdk struct {
	ServerUrl string
	Token     string
}

// Play plays a media file
func (sdk *RcScriptSdk) Play(sessionId, partyId string, body PlayRequest) (*http.Response, error) {
	apiUrl := urlutil.JoinAbsolute(sdk.ServerUrl,
		fmt.Sprintf(UrlTelephonySessionsPartiesPlayFormat, DefaultParamValue, sessionId, partyId))
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
	req.Header.Add(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJsonUtf8)
	return client.Do(req)
}

// Hangsup a call
func (sdk *RcScriptSdk) Hangup(sessionId string) (*http.Response, error) {
	apiUrl := urlutil.JoinAbsolute(sdk.ServerUrl,
		fmt.Sprintf(UrlTelephonySessionFormat, DefaultParamValue, sessionId))
	fmt.Println(apiUrl)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, apiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(httputilmore.HeaderAuthorization, "Bearer "+sdk.Token)
	return client.Do(req)
}
