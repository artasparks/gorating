package gorating

// Representation of the Game
type Game interface {
	// Identifier of player one.
	PlayerOne() Player

	// Identifier of player two.
	PlayerTwo() Player

	// The result of the game. Meant to be from the perspective of Player One.
	GameResult() Result
}

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

type BaseGame struct {
	Player1    Player
	Player2    Player
	GameResult Result
}

type BaseResult struct {
	ResultScore float64
	ResultType  ResultType
}
