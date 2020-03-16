package main

import (
	"net/url"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
	"github.com/maxence-charriere/go-app/v6/pkg/log"
)

type gameSelect struct {
	app.Compo

	filter string
	games  []game
}

func (s *gameSelect) OnNav(u *url.URL) {
	s.showGames()
}

func (s *gameSelect) Render() app.UI {
	return app.Div().
		Class("layout").
		Body(
			app.Div().
				Class("background").
				Body(
					app.Main().Body(
						app.H1().Body(
							app.Text("Select games"),
						),
						app.Input().
							Class("section").
							Type("text").
							Placeholder("Filter").
							Value(s.filter).
							OnKeyup(s.onFilterChange).
							OnChange(s.onFilterChange),
						app.Table().
							Class("section").
							Body(
								app.Range(s.games).Slice(func(i int) app.UI {
									game := s.games[i]

									return app.Tr().
										Class("gamerow").
										DataSet("id", game.ID).
										OnClick(s.onGameSelect).
										Body(
											app.Td().Body(
												app.Div().
													Class("title").
													Body(
														app.Text(game.Name),
													),
												app.Div().
													Class("subtitle").
													Body(
														app.Text(game.Location),
													),
											),
											app.Td().Body(
												app.If(game.Enabled,
													app.Div().
														Class("selected").
														Body(
															app.Text("✓"),
														),
												),
											),
										)
								}),
							),
						app.A().
							Class("app-button section").
							Href("/").
							Body(
								app.Text("Done"),
							),
					),
				),
		)
}

func (s *gameSelect) showGames() {
	games := gameList(myGames)
	games = filterGameList(games, s.filter)
	sortGameList(games)
	s.games = games
	s.Update()
}

func (s *gameSelect) onFilterChange(src app.Value, e app.Event) {
	filter := src.Get("value").String()
	s.filter = filter
	s.Update()

	s.showGames()
}

func (s *gameSelect) onGameSelect(src app.Value, e app.Event) {
	id := src.Get("dataset").Get("id").String()
	game := myGames[id]
	game.Enabled = !game.Enabled
	myGames[id] = game

	if err := saveMyGames(myGames); err != nil {
		log.Error("saving games failed").T("error", err)
	}

	s.Update()
	s.showGames()
}
