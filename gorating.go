// A package to perform two-player based ratings, like elo, for players player
// games.
package gorating

// Representation of a rating.
//
// Most rating systems have various supporting values such as variance,
// deviation, and other such values.
type Rating interface {
	// Gets the current rating, as a numeric rating.
	NumericScore() float64
}

type CalculatedRating interface {
	RatingResult() Rating
	PlayerId() string
}

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
	PlayerOne() CompareablePlayer

	// Retrieves the second player.
	PlayerTwo() CompareablePlayer

	// Retrieves the game result.
	GameResult() Result
}

// A system for rating a player. The Rating System interface is the
// goal for the rating systems defined in the subdirectories.
type RatingSystem interface {
	// Rate all the players who played in a tournament.
	AllPlayersForEvent([]Game) []CalculatedRating

	// Rate only a single player who played in a tournament.
	//
	// Returns nil if the player is not specified in the relevant games.
	PlayerForEvent(CompareablePlayer, []Game) CalculatedRating

	// Rate two players who played a single game.
	BothPlayersInstant(Game) []CalculatedRating

	// Rate two players who played a single game.
	//
	// Returns nil if the player is not specified in the game.
	PlayerInstant(CompareablePlayer, Game) CalculatedRating
}
