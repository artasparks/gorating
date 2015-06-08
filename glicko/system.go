package goglicko

import (
	"fmt"
	gr "github.com/Kashomon/gorating"
)

///////////////////////
// The Rating System //
///////////////////////

// Calculated Rating.
type CalcRating struct {
	R  *Rating
	Id string
}

func (t *CalcRating) Rating() gr.Ratable {
	return gr.Ratable(t.R)
}

func (t *CalcRating) UniqueId() string {
	return t.Id
}

type System struct {
	rmap map[string][]gr.Game

	pmap map[string]gr.PlayerRating
}

func (t *System) RateAll(games []gr.Game) ([]gr.PlayerRating, error) {
	t.rmap, t.pmap = gr.PlayerMaps(games)
	// Rough estimate: assume each player has played only once.
	out := make([]gr.PlayerRating, 0, len(t.rmap))
	for id, games := range t.rmap {
		if _, ok := t.pmap[id]; !ok {
			panic("No Player Rating for id: " + id)
		}
		newr, err := t.Rate(t.pmap[id], games)
		if err != nil {
			return nil, err
		}
		out = append(out, newr)
	}
	return out, nil
}

func (t *System) Rate(player gr.PlayerRating, games []gr.Game) (gr.PlayerRating, error) {
	var filtered []gr.Game
	if t.rmap != nil {
		filtered = t.rmap[player.UniqueId()]
	} else {
		filtered = gr.FilterGames(player, games)
	}
	if len(filtered) == 0 {
		return nil, fmt.Errorf("No games found for player with Id: %s.", player.UniqueId())
	}
	pRate, ok := player.Rating().(*Rating)
	if !ok {
		return nil, fmt.Errorf("Rating %s not a glicko rating for player: %s.", player.Rating(), player.UniqueId())
	}
	otherRates := make([]*Rating, 0, len(games))
	results := make([]Result, 0, len(games))
	for _, g := range filtered {
		var r gr.Ratable
		if p1 := g.PlayerOne(); p1.UniqueId() != player.UniqueId() {
			r = p1.Rating()
		} else if p2 := g.PlayerTwo(); p2.UniqueId() != player.UniqueId() {
			r = p2.Rating()
		} else {
			panic(fmt.Sprintf("Impossible Game: Player one played twice: %v", g))
		}
		glickor, ok := r.(*Rating)
		if !ok {
			return nil, fmt.Errorf("Rating not a Glicko rating for game: %v", g)
		}
		otherRates = append(otherRates, glickor)
		results = append(results, Result(g.GameResult()))
	}
	newRating, err := CalculateRating(pRate, otherRates, results)
	if err != nil {
		return nil, err
	} else {
		return &CalcRating{
			newRating,
			player.UniqueId(),
		}, nil
	}
}
