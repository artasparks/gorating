package gorating

type ResultType int64

const (
	Success ResultType = iota
	Failure
	Anulled
	Forfeit
)

// A game result.
type Result interface {
	Score() float64
	Type() ResultType
}

// Representation of the Game
type Game interface {
	// ID of player two.
	PlayerOneId() string

	// ID of player one.
	PlayerTwoId() string

	// The result of the game.
	Result() Result
}

// Representation of the Tournament
type Tournament interface {
	// Return Tournament games.
	Games() []Game

	// Get the games for a particular Player
	GamesForPlayer(Player) []Game

	// Return the players who played in the tournament.
	Players() []Player
}

// A system for rating a player
type RatingSystem interface {
	// Rate a player, who played in a tournament.
	RateFromTournament(Player, Tournament) Rating

	// Rate a player, who played in a game.
	RateFromGame(Player, Game) Rating
}
