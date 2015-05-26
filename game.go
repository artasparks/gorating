package gorating

// Game result types are meant to encapsulate the idea that games can terminate in
// multiple ways. The most common scenario is that games terminate due to a win
// or a loss. However, games (esp. in tournaments) can also terminate because
// players don't show up, there was a rule confict, or due to myriad of other
// reasons. Annullment and forfeiture are meant to encapsulate these
// conditionsand forfeiture are meant to encapsulate these conditions.
//
// Note also that ResultTypes are always from the perspective of one player
// (usu. player 1) but that's not specified here.
//
//	Win/Loss: Wether or not a specific player won or lost.
//
//	Forfeit(Win/Loss): This is typically equivalent to a win or loss for a
//	specific player, but, unlike a win/loss, it's often the case that a game
//	wasn't played. One common cause of a forfeiture result is that a player
//	doesn't show up.
//
//	Partial: Rating systems may give partial points to either player for a
//	variety of reasons. This is meant to encapsulate such a result.
//
//	Draw: Draw for both players: I.e., neither a win nor a loss.
//
//	Anulled: The game result has been anulled (and so shouldn't affect the
//	players' rankings). Typically this means there was some problem on the part
//	of the tournament director.
type ResultType int64

const (
	Win ResultType = iota
	Loss
	Partial
	Draw
	ForfeitWin
	ForfeitLoss
	Anulled
)

// A game result, I.e., two players and the result of a game (numerical score).
//
//	- Score. Typically this is a numeric score from 0 to 1. However, the actual
//	interpretation of the score is not specified here. Various rating systems
//	are free to use there own result metric. Note that this is always from the
//	perspective of one player.
type Result struct {
	Score float64
}

// An instance of a Game. Two 'players' and the result of their game.
//
// In general, a Game need not represent an actual game and the players need not
// represent actual players. It could represent AI players playing games, or
// even a player attempting a problem.
//
// The Result is from the perspective of the first player.
type Game struct {
	PlayerOne  CompareablePlayer
	PlayerTwo  CompareablePlayer
	GameResult Result
}
