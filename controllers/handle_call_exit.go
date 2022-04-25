package controllers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/net/httputilmore"
	"github.com/grokify/ringcentral-appointment-reminder-demo/rcscript"
	log "github.com/sirupsen/logrus"
)

func HandleCallExit() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("EVT_RECEIVE__ON_CALL_EXIT")
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Warn(err.Error())
		} else {
			log.Info(string(bytes))
		}
	}
}

func (h *Handlers) HandleCommandUpdate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("EVT_RECEIVE__ON_COMMAND_UPDATE")

		var evt rcscript.CommandUpdateEvent
		err := rcscript.Bind(&evt, r)
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(evt)

		w.WriteHeader(http.StatusNoContent)

		if evt.Command == rcscript.CommandPlay && evt.Status == rcscript.StatusCompleted {
			go hangup(h.RcScriptSdk, evt.SessionID)
		}
	}
}

func HandleCommandError() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Warn("EVT_RECEIVE__ON_COMMAND_ERROR: READ_RR: " + err.Error())
		} else {
			log.Warn("EVT_RECEIVE__ON_COMMAND_ERROR: EVT_BODY: " + string(bytes))
		}
	}
}

func hangup(sdk rcscript.RcScriptSdk, telephonySessionID string) {
	time.Sleep(1 * time.Second)

	resp, err := sdk.Hangup(telephonySessionID)
	if err != nil {
		log.Warn(fmt.Sprintf("Play_API_Error: %v\n", err.Error()))
	} else {
		log.Info(fmt.Sprintf("Play_API_Status: %v\n", resp.Status))
	}
	httputilmore.PrintResponse(resp, true)
	log.Info("HANGUP__DONE")
}
