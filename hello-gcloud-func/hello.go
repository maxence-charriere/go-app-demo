package hello

import (
	"net/http"
	"os"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

var handler = &app.Handler{
	Title:   "Hello Google Cloud Function",
	Author:  "Maxence Charriere",
	RootDir: "https://storage.googleapis.com/goapp-gcloud-function",
	Version: os.Getenv("K_REVISION"),
}

// Hello is the Google Cloud Function implementation.
func Hello(w http.ResponseWriter, r *http.Request) {
	handler.ServeHTTP(w, r)
}
