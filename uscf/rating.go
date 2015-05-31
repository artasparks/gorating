//
// The USCF ELO implementation. For more details, see:
//
// http://www.glicko.net/ratings/rating.system.pdf
//
package uscf

type EloRating struct {
	Score float64
	Prov  bool
}

// Construct a new rating
func NewRating() *EloRating {
	return &EloRating{
		750,
		true,
	}
}

func NewRatingEstimated(val float64) *EloRating {
	return &EloRating{
		val,
		false,
	}
}

// Convert from a FIDE ELO Rating.
func FromFide(fide float64) *EloRating {
	converted := 0.0
	if fide <= 2000 {
		converted = fide*0.76 + 560
	} else {
		converted = fide + 80
	}
	return &EloRating{
		converted,
		false,
	}
}

// For compatibility with the ratable interface.
func (t *EloRating) NumericScore() float64 {
	return t.Score
}

// Players playing ol
func (t *EloRating) IsProvisional() bool {
	return t.Prov
}

// Sets the provisional status from the number of games played
func (t *EloRating) SetProvFromGames(games int64) {
	if games < 8 {
		t.Prov = true
	} else {
		t.Prov = false
	}
}
