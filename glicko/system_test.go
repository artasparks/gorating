package goglicko

import (
	gr "github.com/Kashomon/gorating"
	"testing"
)

type player struct {
	r  *Rating
	id string
}

func (t *player) UniqueId() string {
	return t.id
}

func (t *player) Rating() gr.Ratable {
	v := gr.Ratable(t.r)
	return v
}

var _ gr.PlayerRating = &player{}

type tgame struct {
	p1     *player
	p2     *player
	result float64
}

func (t *tgame) PlayerOne() gr.PlayerRating { v := gr.PlayerRating(t.p1); return v }
func (t *tgame) PlayerTwo() gr.PlayerRating { v := gr.PlayerRating(t.p2); return v }
func (t *tgame) GameResult() float64        { return t.result }

func TestFoo(t *testing.T) {
	sys := &System{}

	oldp := &player{
		NewRating(1500, 200, DefaultVol),
		"player0",
	}

	games := []gr.Game{
		&tgame{
			&player{
				NewRating(1400, 30, DefaultVol),
				"player1",
			},
			oldp,
			1,
		},
		&tgame{
			&player{
				NewRating(1550, 100, DefaultVol),
				"player2",
			},
			oldp,
			0,
		},
		&tgame{
			&player{
				NewRating(1700, 300, DefaultVol),
				"player2",
			},
			oldp,
			0,
		},
	}

	newval, err := sys.Rate(oldp, games)
	if err != nil {
		t.Errorf("err: %s", err)
	}
	if !floatsMostlyEqual(newval.Rating().NumericScore(), 1464.06, 0.01) {
		t.Errorf("Unexpected rating: %v", newval)
	}
}
