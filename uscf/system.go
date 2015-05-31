package uscf

import gr "github.com/Kashomon/gorating"

//
// The USCF Rating system
//

// The provisional winning expectency.
//
// Basically, if the delta between the player and opponent's score is greater
// than 400, we consider a win to be guaranteed.
func provWinExpect(player, opp *EloRating) float64 {
	delta := player.Score - opp.Score
	if delta <= -400 {
		return 0
	} else if delta >= 400 {
		return 1
	} else {
		return 0.5 + delta/800.0
	}
}

type UscfSystem struct {
	// Map from player id to games played
	pmap map[string][]gr.Game
}

func (t *UscfSystem) AllPlayersForEvent(games []gr.Game) []*gr.CalcRating {
	t.pmap = gr.PlayerMap(games)
	return nil
}

func (t *UscfSystem) PlayerForEvent(player gr.Player, games []gr.Game) []*gr.CalcRating {
	t.pmap = gr.PlayerMap(games)

	_, ok := t.pmap[player.UnqiueId()]
	if !ok {
		return []*gr.CalcRating{}
	}
	return nil
}

var _ gr.RatingSystem = &UscfSystem{}
