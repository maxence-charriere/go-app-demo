package hello

import (
	"net/http"
	"os"

	"github.com/maxence-charriere/go-app/pkg/app"
)

var handler = &app.Handler{
	Title:   "Hello App Engine",
	Author:  "Maxence Charriere",
	Version: os.Getenv("K_REVISION"),
}

// Hello is the Google Cloud Function implementation.
func Hello(w http.ResponseWriter, r *http.Request) {
	handler.ServeHTTP(w, r)
}
