package main

import (
	"net/url"
	"strconv"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
	"github.com/maxence-charriere/go-app/v6/pkg/log"
)

type home struct {
	app.Compo

	luckyNumber int
	games       []game
}

func (h *home) OnNav(u *url.URL) {
	if err := app.LocalStorage.Get("luckyNumber", &h.luckyNumber); err != nil {
		log.Error("retrieving lucky number failed").T("error", err)
	}

	games := gameList(myGames)
	sortGameList(games)
	h.games = games
	h.Update()
}

func (h *home) Render() app.UI {
	return app.Div().
		Class("layout").
		Body(
			app.Div().
				Class("background").
				Body(
					app.Main().Body(
						app.H1().Body(
							app.Text("What is your lucky number?"),
						),
						app.Input().
							Class("section").
							Type("number").
							Placeholder("Lucky number").
							Value(h.luckyNumber).
							OnChange(h.onLuckyNumberChange).
							OnKeyup(h.onLuckyNumberChange),
						app.Table().
							Class("section").
							Body(
								app.Range(h.games).Slice(func(i int) app.UI {
									game := h.games[i]
									if !game.Enabled {
										return nil
									}

									return app.Tr().Body(
										app.Td().Body(
											app.Div().Body(app.Text(game.Name)),
											app.Div().
												Class("subtitle accent").
												Body(app.Text(game.Location)),
										),
										app.Td().
											Class("draw").
											Body(
												app.Text(game.draw(h.luckyNumber)),
											),
									)
								}),
							),
						app.A().
							Class("app-button section").
							Href("select").
							Body(
								app.Text("Select other games"),
							),
					),
				),
		)
}

func (h *home) onLuckyNumberChange(src app.Value, e app.Event) {
	number, _ := strconv.Atoi(src.Get("value").String())
	h.luckyNumber = number
	h.Update()

	if err := app.LocalStorage.Set("luckyNumber", number); err != nil {
		log.Error("saving lucky number failed").T("error", err)
	}
}
