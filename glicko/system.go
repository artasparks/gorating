package goglicko

import (
	gr "github.com/Kashomon/gorating"
)

///////////////////////
// The Rating System //
///////////////////////

type System struct {
}

func (t *System) RateAll(games []gr.Game) ([]gr.PlayerRating, error) {
	// t.pmap = gr.PlayerMap(games)
	return nil, nil
}

func (t *System) Rate(player gr.Player, games []gr.Game) (gr.PlayerRating, error) {
	return nil, nil
}

var _ gr.RatingSystem = &System{}
