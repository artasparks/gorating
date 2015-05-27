//
// A simple rating implementation
//
package simple

////////////////////
// Example Player //
////////////////////

// A simple player instance. Most practical systems will want to use more complex
// players. However, this can be useful for scripts. or for tests.
type SimplePlayer struct {
	// ID of the player.
	Id string

	// Rating of the player.
	R Rating
}

// Creates a new base player.
func NewBasePlayer(id string) *BasePlayer {
	return &BasePlayer{
		id,
		DefaultBaseRating(),
	}
}

// Get the UnqiueId, to make this a CompareablePlayer.
func (t *BasePlayer) UnqiueId() string {
	return t.Id
}

// Get the PlayerRating, to make this a CompareablePlayer.
func (t *BasePlayer) PlayerRating() Rating {
	return t.R
}

////////////////////////////
// Example implementation //
////////////////////////////

// A simple rating example
type SimpleRating struct {
	score float64
	prov  bool
}

func (t *SimpleRating) NumericScore() float64 {
	return t.score
}

func (t *SimpleRating) IsProvisional() bool {
	return t.prov
}
