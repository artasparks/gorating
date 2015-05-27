//
// A simple rating implementation
//
package simple

import "github.com/Kashomon/gorating"

// A simple player instance. Most practical systems will want to use more complex
// players. However, this can be useful for scripts. or for tests.
type SimplePlayer struct {
	// ID of the player.
	id string

	// Rating of the player.
	score float64
}

// Get the UnqiueId, to make this a CompareablePlayer.
func (t *SimplePlayer) UnqiueId() string {
	return t.id
}

// Get the PlayerRating, to make this a CompareablePlayer.
func (t *SimplePlayer) NumericScore() float64 {
	return t.score
}

// Ensure that the Simple player satisfies the Player interface
var _ gorating.Player = &SimplePlayer{}

type SimpleGame struct {
	p1     *SimplePlayer
	p2     *SimplePlayer
	result float64
}

func (t *SimpleGame) PlayerOne() gorating.Player {
	return gorating.Player(t.p1)
}

func (t *SimpleGame) PlayerTwo() gorating.Player {
	return gorating.Player(t.p2)
}

func (t *SimpleGame) GameResult() float64 {
	return t.result
}

var _ gorating.Game = &SimpleGame{}

type SimpleRatingSystem struct {
}

func (t *SimpleRatingSystem) AllPlayersForEvent([]gorating.Game) []gorating.Player {
	return []gorating.Player{}
}

func (t *SimpleRatingSystem) PlayerForEvent(gorating.Player, []gorating.Game) gorating.Player {
	return gorating.Player(&SimplePlayer{})
}

var _ gorating.RatingSystem = &SimpleRatingSystem{}
