package controllers

import (
	"fmt"
	"net/http"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/ringcentral-appointment-reminder-demo/rcscript"
	log "github.com/sirupsen/logrus"
)

const UrlStarWarsMainTheme string = "https://www.thesoundarchive.com/starwars/star-wars-theme-song.mp3"

type Handlers struct {
	RcScriptSdk rcscript.RcScriptSdk
}

func (h *Handlers) HandleCallEnter() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ON_CALL_ENTER")

		var evt rcscript.CallEnterEvent
		err := rcscript.Bind(&evt, r)
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(evt)

		play := rcscript.PlayRequest{
			Resources: []rcscript.Resource{
				rcscript.Resource{Uri: UrlStarWarsMainTheme},
			},
			InterruptByDtmf: false,
			RepeatCount:     1}
		fmtutil.PrintJSON(play)

		resp, err := h.RcScriptSdk.Play(evt.SessionId, evt.InParty.Id, play)
		if err != nil {
			log.Warn(fmt.Sprintf("Play_API_Error: %v\n", err.Error()))
		} else {
			log.Info(fmt.Sprintf("Play_API_Status: %v\n", resp.Status))
		}
		httputilmore.PrintResponse(resp, true)

		fmt.Fprintf(w, "<!DOCTYPE html><html><body><h1>HandleOnCallEnter</h1></body></html>")
	}
}
