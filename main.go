package main

import (
	"net/http"
	"os"

	"github.com/grokify/ringcentral-appointment-reminder-demo/controllers"
	"github.com/grokify/ringcentral-appointment-reminder-demo/rcscript"
	"github.com/grokify/simplego/config"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
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

	log.Info("Listening on phone number: " + os.Getenv("APP_NUMBER"))

	sdk := rcscript.RcScriptSdk{
		ServerUrl: os.Getenv("RINGCENTRAL_SERVER_URL"),
		Token:     os.Getenv("RINGCENTRAL_ACCESS_TOKEN")}
	if len(sdk.ServerUrl) == 0 {
		log.Fatal("E_INIT_FAILURE__NO_RINGCENTRAL_SERVER_URL")
	} else if len(sdk.Token) == 0 {
		log.Fatal("E_INIT_FAILURE__NO_RINGCENTRAL_ACCESS_TOKEN")
	}

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

	portStr := ":" + DefaultPort
	port := os.Getenv("PORT")
	if len(port) > 0 {
		portStr = ":" + port
	}
	log.Infof("Running on [%v]\n", portStr)
	http.ListenAndServe(portStr, nil)
}
