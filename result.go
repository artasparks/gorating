// A package to perform two-player based ratings, like elo, for players player
// games.
package gorating

// Game result types are meant to encapsulate the idea that games can terminate
// in multiple ways, but that the results can usually be binned into one of a
// few conditions. The most common scenario is that games terminate due to a win
// or a loss. However, games (esp. in tournaments) can also terminate because
// players don't show up, there was a rule confict, or due to myriad of other
// reasons.
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

// Convert a ResultType to a Result.
func ResultForResultType(r ResultType) float64 {
	switch r {
	case Win:
	case ForfeitWin:
		return 1.0
	case ForfeitLoss:
	case Loss:
		return 0.0
	case Anulled:
	case Draw:
		return 0.5
	case Partial:
	default:
		return 0.5
	}
	return 0.5
}
