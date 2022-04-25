package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/net/httputilmore"
	"github.com/grokify/ringcentral-appointment-reminder-demo/rcscript"
	log "github.com/sirupsen/logrus"
)

const (
	MediaURLStarWarsMainTheme         string = "https://www.thesoundarchive.com/starwars/star-wars-theme-song.mp3"
	MediaURLStarWarsMainThemeDecodedG string = "http://raw.githubusercontent.com/vshisterov/test/master/star_wars_decoded.wav"
	MediaURLStarWarsMainThemeDecoded  string = "http://ringforce.org/star-wars_decoded.wav"
	MediaURLThankYouRc                string = "http://ringforce.org/thank-you-for-using-rc.wav"
	MediaURLThankYouRc1               string = "http://10.28.21.3/prompts/english__united_states_/thankyouforusingrc.wav"
	MediaURLThankYouRc2               string = "http://raw.githubusercontent.com/vshisterov/test/master/thankyouforusingrc.wav"
	MediaURLStarWarsWav               string = "http://www.moviewavs.com/0053148414/WAVS/Movies/Star_Wars/starwars.wav"
)

type Handlers struct {
	RcScriptSdk rcscript.RcScriptSdk
}

func play(sdk rcscript.RcScriptSdk, evt rcscript.CallEnterEvent) {
	fmtutil.MustPrintJSON(evt)
	time.Sleep(1 * time.Second)
	play := rcscript.PlayRequest{
		Resources: []rcscript.Resource{
			{URI: MediaURLStarWarsMainThemeDecoded}},
		InterruptByDtmf: false,
		RepeatCount:     1}
	fmtutil.MustPrintJSON(play)

	resp, err := sdk.Play(evt.SessionID, evt.InParty.ID, play)
	if err != nil {
		log.Warn(fmt.Sprintf("Play_API_Error: Status [%v] Message[%v]\n", resp.Status, err.Error()))
	} else {
		log.Info(fmt.Sprintf("Play_API_Status: %v\n", resp.Status))
	}
	logutil.FatalErr(httputilmore.PrintResponse(resp, true))
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
