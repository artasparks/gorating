package gorating

// Representation of the Player. ID is meant to be unique
type Player interface {
	// A unique identifier for the Player.
	UnqiueId() string

	// Return the player's rating.
	PlayerRating() Rating
}

type BasePlayer struct {
	// ID of the player.
	Identifier string

	// Rating of the player.
	Rate Rating
}

func DefaultBasePlayer(id string) *BasePlayer {
	return &BasePlayer{
		id,
		DefaultBaseRating(),
	}
}

func (t *BasePlayer) UnqiueId() string {
	return t.Identifier
}

func (t *BasePlayer) PlayerRating() Rating {
	return t.Rate
}
