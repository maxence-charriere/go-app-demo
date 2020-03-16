package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
	"github.com/maxence-charriere/go-app/v6/pkg/log"
)

type game struct {
	ID        string
	Name      string
	Location  string
	Enabled   bool
	NotSorted bool
	Numbers   []numberConfig
}

func (g game) draw(n int) string {
	r := rand.New(rand.NewSource(int64(n)))
	w := strings.Builder{}

	for i, cfg := range g.Numbers {
		if i != 0 {
			fmt.Fprint(&w, " - ")
		}

		numbers := make([]int, 0, cfg.Count)
		cache := make(map[int]struct{}, cfg.Count)

		for len(numbers) < cfg.Count {
			n := r.Intn(cfg.Max) + 1
			if _, ok := cache[n]; ok {
				continue
			}
			cache[n] = struct{}{}
			numbers = append(numbers, n)
		}

		sort.Ints(numbers)

		for _, n := range numbers {
			if n < 10 {
				fmt.Fprintf(&w, "0%v ", n)
			} else {
				fmt.Fprintf(&w, "%v ", n)
			}
		}
	}

	return strings.TrimSpace(w.String())
}

type numberConfig struct {
	Count int
	Max   int
}

var (
	games = map[string]game{
		"powerball": {
			ID:       "powerball",
			Name:     "Powerball",
			Location: "USA",
			Enabled:  true,
			Numbers: []numberConfig{
				{
					Count: 5,
					Max:   69,
				},
				{
					Count: 1,
					Max:   26,
				},
			},
		},
		"megamillions": {
			ID:       "megamillions",
			Name:     "MEGA Millions",
			Location: "USA",
			Numbers: []numberConfig{
				{
					Count: 5,
					Max:   70,
				},
				{
					Count: 1,
					Max:   25,
				},
			},
		},
		"superlottoplus": {
			ID:       "superlottoplus",
			Name:     "SuperLotto Plus",
			Location: "USA",
			Numbers: []numberConfig{
				{
					Count: 5,
					Max:   47,
				},
				{
					Count: 1,
					Max:   27,
				},
			},
		},
		"fantasy5": {
			ID:       "fantasy5",
			Name:     "Fantasy 5",
			Location: "USA",
			Numbers: []numberConfig{
				{
					Count: 5,
					Max:   39,
				},
			},
		},
		"daily3": {
			ID:        "daily3",
			Name:      "Daily 3",
			Location:  "USA",
			NotSorted: true,
			Numbers: []numberConfig{
				{
					Count: 3,
					Max:   9,
				},
			},
		},
		"daily4": {
			ID:        "daily4",
			Name:      "Daily 4",
			Location:  "USA",
			NotSorted: true,
			Numbers: []numberConfig{
				{
					Count: 4,
					Max:   9,
				},
			},
		},
		"euromillions": {
			ID:       "euromillions",
			Name:     "EuroMillions",
			Location: "Europe",
			Enabled:  true,
			Numbers: []numberConfig{
				{
					Count: 5,
					Max:   50,
				},
				{
					Count: 2,
					Max:   12,
				},
			},
		},
		"loto": {
			ID:       "loto",
			Name:     "Loto",
			Location: "France",
			Numbers: []numberConfig{
				{
					Count: 5,
					Max:   49,
				},
				{
					Count: 1,
					Max:   10,
				},
			},
		},
		"loto6": {
			ID:       "loto6",
			Name:     "Loto 6",
			Location: "Japan",
			Enabled:  true,
			Numbers: []numberConfig{
				{
					Count: 6,
					Max:   43,
				},
			},
		},
		"loto7": {
			ID:       "loto7",
			Name:     "Loto 7",
			Location: "Japan",
			Numbers: []numberConfig{
				{
					Count: 7,
					Max:   37,
				},
			},
		},
	}

	myGames map[string]game
)

func init() {
	g, err := loadMyGames()
	if err != nil {
		g = games
		log.Error("loading my games failed").T("error", err)
	}

	myGames = g
}

func loadMyGames() (map[string]game, error) {
	myGames := make(map[string]game, len(games))
	if err := app.LocalStorage.Get("myGames", &games); err != nil {
		return nil, err
	}

	for k := range myGames {
		if _, ok := games[k]; !ok {
			delete(myGames, k)
		}
	}

	for k, v := range games {
		if old, ok := myGames[k]; ok {
			v.Enabled = old.Enabled
		}
		myGames[k] = v
	}

	return myGames, nil
}

func saveMyGames(g map[string]game) error {
	return app.LocalStorage.Set("myGames", g)
}

func gameList(games map[string]game) []game {
	l := make([]game, 0, len(games))
	for _, g := range games {
		l = append(l, g)
	}
	return l
}

func sortGameList(games []game) {
	sort.Slice(games, func(i, j int) bool {
		return strings.Compare(games[i].Name, games[j].Name) <= 0
	})
}

func filterGameList(games []game, filter string) []game {
	var f []game
	for _, g := range games {
		if filter == "" ||
			strings.Contains(
				strings.ToLower(g.Name),
				strings.ToLower(filter),
			) ||
			strings.Contains(
				strings.ToLower(g.Location),
				strings.ToLower(filter),
			) {
			f = append(f, g)
		}
	}
	return f
}
