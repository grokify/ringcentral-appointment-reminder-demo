package rcscript

type PlayRequest struct {
	Resources       []Resource `json:"resources"`
	InterruptByDtmf bool       `json:"interruptByDtmf"`
	RepeatCount     int        `json:"repeatCount"`
}

type Resource struct {
	Uri string `json:"uri"`
}

/*

{
	"resources": [
	  {
		"uri": "http://example.com/ivr-app-example/greeting.wav"
  } ],
	"interruptByDtmf": false,
	"repeatCount": 1
  }
*/
