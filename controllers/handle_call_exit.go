package controllers

import (
	"net/http"
)

func HandleCallExit() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func HandleCommandUpdate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func HandleCommandError() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
