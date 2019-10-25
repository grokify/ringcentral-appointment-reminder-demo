package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/ringcentral-appointment-reminder-demo/controllers"
	"github.com/grokify/ringcentral-appointment-reminder-demo/rcscript"
	"github.com/jessevdk/go-flags"
)

const DefaultPort string = "8080"

type Options struct {
	EnvFile string `short:"e" long:"env" description:"Env filepath"`
}

func setup() controllers.Handlers {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}
	if len(opts.EnvFile) > 0 {
		err := config.LoadDotEnvSkipEmpty(opts.EnvFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	sdk := rcscript.RcScriptSdk{
		ServerUrl: os.Getenv("RINGCENTRAL_SERVER_URL"),
		Token:     os.Getenv("APP_ACCESS_TOKEN")}

	fmtutil.PrintJSON(sdk)
	handlers := controllers.Handlers{
		RcScriptSdk: sdk}
	return handlers
}

func main() {
	handlers := setup()

	http.HandleFunc("/on-call-enter", handlers.HandleCallEnter())
	http.HandleFunc("/on-call-exit", controllers.HandleCallExit())
	http.HandleFunc("/on-command-update", handlers.HandleCommandUpdate())
	http.HandleFunc("/on-command-error", controllers.HandleCommandError())
	http.HandleFunc("/ping", controllers.HandlePing())

	port := os.Getenv("PORT")
	portStr := ":" + DefaultPort
	if len(port) > 0 {
		portStr = ":" + port
	}
	fmt.Printf("Running on [%v]\n", portStr)
	http.ListenAndServe(portStr, nil)

}
