package controllers

import (
	"fmt"
	"net/http"
)

func HandlePing() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<!DOCTYPE html><html><body><h1>Pong</h1></body></html>")
	}
}
