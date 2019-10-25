package rcscript

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/gotilla/net/urlutil"
)

const playUrlFormat string = `/restapi/v1.0/account/~/telephony/sessions/%s/parties/%s/play`

// RcScriptSdk is a simple SDK for making Call Scripting Commands
type RcScriptSdk struct {
	ServerUrl string
	Token     string
}

// Play plays a media file
func (sdk *RcScriptSdk) Play(sessionId, partyId string, body PlayRequest) (*http.Response, error) {
	apiUrl := urlutil.JoinAbsolute(sdk.ServerUrl,
		fmt.Sprintf(playUrlFormat, sessionId, partyId))
	fmt.Println(apiUrl)

	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	fmt.Printf("TOKEN: %s\n", sdk.Token)
	req.Header.Add(httputilmore.HeaderAuthorization, "Bearer "+sdk.Token)
	req.Header.Add(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJsonUtf8)
	return client.Do(req)
}
