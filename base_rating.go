package gorating

const (
	DefaultRat = 1500.0 // Default starting rating
	DefaultDev = 350.0  // Default starting deviation
	DefaultVol = 0.06   // Default starting volatility
)

// A default rating struct
type BaseRating struct {
	// Rating Number
	rating float64

	// The Deviation
	deviation float64

	// The Volatility
	volatility float64
}

// Construct a new default rating
func DefaultBaseRating() *BaseRating {
	return &BaseRating{
		DefaultRat,
		DefaultDev,
		DefaultVol,
	}
}

func (r *BaseRating) Rating() float64 {
	return r.rating
}
