package controllers

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/grokify/mogo/net/http/httputilmore"
	"github.com/grokify/ringcentral-appointment-reminder-demo/rcscript"
)

func HandleCallExit() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("EVT_RECEIVE__ON_CALL_EXIT")
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Print("WARN: " + err.Error())
		} else {
			log.Print(string(bytes))
		}
	}
}

func (h *Handlers) HandleCommandUpdate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("EVT_RECEIVE__ON_COMMAND_UPDATE")

		var evt rcscript.CommandUpdateEvent
		err := rcscript.Bind(&evt, r)
		if err != nil {
			log.Fatal(err)
		}
		// fmtutil.PrintJSON(evt)

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
			log.Print("EVT_RECEIVE__ON_COMMAND_ERROR: READ_RR: " + err.Error())
		} else {
			log.Print("EVT_RECEIVE__ON_COMMAND_ERROR: EVT_BODY: " + string(bytes))
		}
	}
}

func hangup(sdk rcscript.RcScriptSdk, telephonySessionID string) {
	time.Sleep(1 * time.Second)

	resp, err := sdk.Hangup(telephonySessionID)
	if err != nil {
		log.Printf("Play_API_Error: %v\n", err.Error())
	} else {
		log.Printf("Play_API_Status: %v\n", resp.Status)
	}
	err = httputilmore.PrintResponse(resp, true)
	if err != nil {
		log.Printf("Play_API_Print_Error: %v\n", err.Error())
	}
	log.Print("HANGUP__DONE")
}
