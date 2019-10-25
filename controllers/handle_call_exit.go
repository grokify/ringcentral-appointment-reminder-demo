package controllers

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func HandleCallExit() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("EVT_RECEIVE__ON_CALL_EXIT")
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Warn(err.Error())
		} else {
			log.Info(string(bytes))
		}
	}
}

func HandleCommandUpdate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("EVT_RECEIVE__ON_COMMAND_UPDATE")
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Warn(err.Error())
		} else {
			log.Info(string(bytes))
		}
	}
}

func HandleCommandError() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("EVT_RECEIVE__ON_COMMAND_ERROR")
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Warn(err.Error())
		} else {
			log.Info(string(bytes))
		}
	}
}
