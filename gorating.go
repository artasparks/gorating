// A package to perform two-player based ratings, like elo, for players player
// games.
package gorating

// A player.  All that we care about is that the player has a unique ID.
type Player interface {
	// Gets the player's un
	UniqueId() string
}

// Ratable interface.
//
// Most rating systems have various supporting values such as variance,
// deviation, and other such values.
type Ratable interface {
	// Gets the current rating, as a numeric rating. This will have wildly
	// different interpretations based on the various scoring systems, but
	// generally, all rating systems return some sort of number at the end of the
	// day.
	NumericScore() float64
}

// Necessary methods so that we can compare players.
//
// RatablePlayer instances must implement
// - UniqueId: A method to retrieve a unique ID.
// - Rating: A way to get the current player's rating
type PlayerRating interface {
	Player
	Rating() Ratable
}

// An instance of a Game. Two 'players' and the result of their game.
//
// In general, a Game need not represent an actual game and the players need not
// represent actual players. It could represent AI players playing games, or
// even a player attempting a problem.
//
// The Result is from the perspective of the first player.
type Game interface {
	// Retrieves the first player. The result should be from the perspective of
	// this player.
	PlayerOne() PlayerRating

	// Retrieves the second player.
	PlayerTwo() PlayerRating

	// A game result that indicates what happ
	//
	// Typically this is a numeric score from 0 to 1. However, the actual
	// interpretation of the score is not specified here. Various rating systems
	// are free to use there own result metric. Note that this is always from the
	// perspective of a the first player.
	GameResult() float64
}

// A system for raing a player. The Rating System interface is the
// goal for the rating systems defined in the subdirectories.
//
// It is assumed the the rating system has been constructed so that it knows
// about all the games played.
type RatingSystem interface {
	// Rate all the players who played in a tournament.
	//
	// Note that rating a single game (instant rating) is a special case.
	RateAll([]Game) ([]PlayerRating, error)

	// Rate only a single player who played in a tournament.
	//
	// As with the above, if you need to rate a single game, just pass in a single
	// game.
	//
	// Returns nil if the player is not specified in the relevant games.
	Rate(PlayerRating, []Game) (PlayerRating, error)
}
