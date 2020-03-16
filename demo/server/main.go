// +build !wasm

package main

import (
	"net/http"
	"os"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
	"github.com/maxence-charriere/go-app/v6/pkg/log"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "7777"
	}

	version := os.Getenv("GAE_VERSION")

	log.Info("starting server").
		T("port", port).
		T("version", version)

	err := http.ListenAndServe(":"+port, &app.Handler{
		Title: "App Demo",
		Styles: []string{
			"/web/hello.css",
			"/web/city.css",
		},
		Version: version,
	})

	if err != nil {
		log.Error("server crashed").T("error", err)
	}
}
