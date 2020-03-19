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
		port = "9000"
	}
	addr := ":" + port

	version := os.Getenv("GAE_VERSION")

	log.Info("starting server").T("addr", addr)

	if err := http.ListenAndServe(addr, &app.Handler{
		Name:        "Luck",
		Author:      "Maxence Charriere",
		Description: "Lottery numbers generator.",
		Icon: app.Icon{
			Default: "/web/icon.png",
		},
		Keywords: []string{
			"lottery",
			"EuroMillions",
			"Loto",
			"MEGA Millions",
			"Powerball",
			"SuperLotto Plus",
		},
		ThemeColor:      "#000000",
		BackgroundColor: "#000000",
		Styles: []string{
			"/web/luck.css",
		},
		CacheableResources: []string{
			"/web/bg.jpg",
		},
		Version: version,
	}); err != nil {
		log.Error("listening and serving http requests failed").
			T("reason", err).
			T("addr", addr).
			T("version", version).
			Panic()
	}
}
