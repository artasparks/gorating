//
// An ELO implementation. For more details, see:
//
// http://www.glicko.net/ratings/rating.system.pdf
//
package elo

type Player struct {
	Id string
}

func NewEloPlayer(id string) *Player {
	return &Player{
		id,
	}
}
