package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/ringcentral-appointment-reminder-demo/rcscript"
	log "github.com/sirupsen/logrus"
)

const (
	MediaUrlStarWarsMainTheme        string = "https://www.thesoundarchive.com/starwars/star-wars-theme-song.mp3"
	MediaUrlStarWarsMainThemeDecoded string = "http://raw.githubusercontent.com/vshisterov/test/master/star_wars_decoded.wav"
	MediaUrlThankYouRc               string = "http://10.28.21.3/prompts/english__united_states_/thankyouforusingrc.wav"
	MediaUrlThankYouRc2              string = "http://raw.githubusercontent.com/vshisterov/test/master/thankyouforusingrc.wav"
)

type Handlers struct {
	RcScriptSdk rcscript.RcScriptSdk
}

func play(sdk rcscript.RcScriptSdk, evt rcscript.CallEnterEvent) {
	time.Sleep(1 * time.Second)
	play := rcscript.PlayRequest{
		Resources: []rcscript.Resource{
			{Uri: MediaUrlStarWarsMainThemeDecoded}},
		InterruptByDtmf: false,
		RepeatCount:     1}
	fmtutil.PrintJSON(play)

	resp, err := sdk.Play(evt.SessionId, evt.InParty.Id, play)
	if err != nil {
		log.Warn(fmt.Sprintf("Play_API_Error: Status [%v] Message[%v]\n", resp.Status, err.Error()))
	} else {
		log.Info(fmt.Sprintf("Play_API_Status: %v\n", resp.Status))
	}
	httputilmore.PrintResponse(resp, true)
	log.Info("PLAY__DONE")
}

func (h *Handlers) HandleCallEnter() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("ON_CALL_ENTER")
		var evt rcscript.CallEnterEvent
		err := rcscript.Bind(&evt, r)
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(evt)

		w.WriteHeader(http.StatusNoContent)

		go play(h.RcScriptSdk, evt)
	}
}
