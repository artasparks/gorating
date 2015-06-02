package goglicko

import (
	"github.com/Kashomon/gorating"
	"testing"
)

type tgame struct {
}

func (t *tgame) PlayerOne() gorating.PlayerRating { return nil }
func (t *tgame) PlayerTwo() gorating.PlayerRating { return nil }
func (t *tgame) GameResult() float64              { return 0 }

func TestFoo(t *testing.T) {
	_ = &System{}
	_ = NewRating(1500, 200, DefaultVol)
	_ = []gorating.Game{
		&tgame{},
	}
}
