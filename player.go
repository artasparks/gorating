package gorating

// Necessary methods so that we can compare players.
//
// CompareablePlayer instances must implement
// - UniqueId: A method to retrieve a unique ID.
// - Rating: A way to get the current player's rating
type CompareablePlayer interface {
	// Gets the unique identifier for the player.
	UnqiueId() string

	// Returns the player's rating.
	PlayerRating() Rating
}

// A basic player instance. Most practical systems will want to use more complex
// players. However, this can be useful for scripts. or for tests.
type BasePlayer struct {
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
