// Implementation of the Glicko 2 Rating system, for rating players.  Glicko is
// an improvoment on ELO, but is much more computationally intensive.
//
// For more information, see:
//
// http://www.glicko.net/glicko/glicko2.pdf
//
// http://en.wikipedia.org/wiki/Glicko_rating_system
//
// The calculation process is broken into 8 steps.
//
// Step 1:
// Determine initial values.
//
// Step 2:
// Convert to Glicko2 Scale from the Glicko1 scale.
//
// Step 3:
// Compute (v), the estimated variance based only on game outcomes.
//
// Step 4:
// Compute the quantity Delta, the estimated improvement.
//
// Step 5:
// Determine the new value, sigma', of the volatility, in an iterative process.
//
// Step 6:
// Update the rating deviation to the new pre-rating period value, φ_z
//
// Step 7:
// Update the rating and RD to the new values, μ′ and φ′:
//
// Step 8:
// Convert back to the Glicko1 scale.
package goglicko

import (
	"fmt"
	"math"
)

// Overrideable Defaults
var (
	// Constrains the volatility. Typically set between 0.3 and 1.2.  Often
	// refered to as the 'system' constant.
	DefaultTau = 0.3

	DefaultRat = 1500.0 // Default starting rating
	DefaultDev = 350.0  // Default starting deviation
	DefaultVol = 0.06   // Default starting volatility
)

// Miscellaneous Mathematical constants.
const (
	piSq = math.Pi * math.Pi // π^2
	// Constant transformation value, to transform between Glicko 2 and Glicko 1
	glicko2Scale = 173.7178
)

// Used to indicate who won/lost/tied the game.
type Result float64

const (
	Win  Result = 1
	Loss Result = 0
	Draw Result = 0.5
)

////////////////////////////
// Sundry of Helper Funcs //
////////////////////////////

// Ensure that two floats are equal, given some epsilon.
func floatsMostlyEqual(v1, v2, epsilon float64) bool {
	return math.Abs(v1-v2) < epsilon
}

// Square function for convenience
func sq(x float64) float64 {
	return x * x
}

// The E function. Written as E(μ,μ_j,φ_j).
// For readability, instead of greek we use the variables
// 	r: rating of player
// 	ri: rating of opponent
// 	devi: deviation of opponent
func ee(r, ri, devi float64) float64 {
	return 1.0 / (1 + math.Exp(-gee(devi)*(r-ri)))
}

// The g function. Written as g(φ).
// For readability, instead of greek we use the variables
// 	dev: The deviation of a player's rating
func gee(dev float64) float64 {
	return 1 / math.Sqrt(1+3*dev*dev/piSq)
}

// Estimate the variance of the team/player's rating based only on game
// outcomes. Note, it must be true that len(ees) == len(gees).
func estVariance(gees, ees []float64) float64 {
	out := 0.0
	for i := range gees {
		out += sq(gees[i]) * ees[i] * (1 - ees[i])
	}
	return 1.0 / out
}

// Estimate the improvement in rating by comparing the pre-period rating to the
// performance rating, based only on game outcomes.
//
// Note: This function is like the 'delta' in the algorithm, but here we don't
// multiply by the estimated variance.
func estImprovePartial(gees, ees []float64, r []Result) float64 {
	out := 0.0
	for i := range gees {
		out += gees[i] * (float64(r[i]) - ees[i])
	}
	return out
}

// Calculate the new volatility for a Player.
func newVolatility(estVar, estImp float64, p *Rating) float64 {
	epsilon := 0.000001
	a := math.Log(sq(p.Volatility))
	deltaSq := sq(estImp)
	phiSq := sq(p.Deviation)
	tauSq := sq(DefaultTau)
	maxIter := 100

	f := func(x float64) float64 {
		eX := math.Exp(x)
		return eX*(deltaSq-phiSq-estVar-eX)/
			(2*sq(phiSq+estVar+eX)) - (x-a)/tauSq
	}

	A := a
	B := 0.0
	if deltaSq > (phiSq + estVar) {
		B = math.Log(deltaSq - phiSq - estVar)
	} else {
		val := -1.0
		k := 1
		for ; val < 0; k++ {
			val = f(a - float64(k)*DefaultTau)
		}
		B = a - float64(k)*DefaultTau
	}
	// Now: A < ln(sigma'^2) < B

	fA := f(A)
	fB := f(B)
	fC := 0.0
	iter := 0
	for math.Abs(B-A) > epsilon && iter < maxIter {
		C := A + (A-B)*fA/(fB-fA)
		fC = f(C)
		if fC*fB < 0 {
			A = B
			fA = fB
		} else {
			fA = fA / 2
		}
		B = C
		fB = fC
		iter++
	}
	if iter == maxIter-1 {
		fmt.Errorf("Fall through! Too many iterations")
	}

	newVol := math.Exp(A / 2)
	return newVol
}

// Calculate the new Deviation.  This is just the L2-norm of the deviation and
// the volatility.
func newDeviation(dev, newVol, estVar float64) float64 {
	phip := math.Sqrt(dev*dev + newVol*newVol)
	return 1.0 / math.Sqrt(1.0/(phip*phip)+1.0/(estVar))
}

// Calculate the new Rating.
func newRatingVal(oldRating, newDev, estImpPart float64) float64 {
	return oldRating + newDev*newDev*estImpPart
}

func CalculateRating(player *Rating, opponents []*Rating, res []Result) (*Rating, error) {
	if len(opponents) != len(res) {
		return nil, fmt.Errorf("Number of opponents must == number of results. %v != %v",
			len(opponents), len(res))
	}

	p2 := player.ToGlicko2()
	gees := make([]float64, len(opponents))
	ees := make([]float64, len(opponents))
	for i := range opponents {
		o := opponents[i].ToGlicko2()
		gees[i] = gee(o.Deviation)
		ees[i] = ee(p2.Rating, o.Rating, o.Deviation)
	}

	estVar := estVariance(gees, ees)
	estImpPart := estImprovePartial(gees, ees, res)
	estImp := estVar * estImpPart
	newVol := newVolatility(estVar, estImp, p2)
	newDev := newDeviation(p2.Deviation, newVol, estVar)
	newRating := newRatingVal(p2.Rating, newDev, estImpPart)
	rt := NewRating(newRating, newDev, newVol).FromGlicko2()

	// Upper bound by the Default Deviation.
	if rt.Deviation > DefaultDev {
		rt.Deviation = DefaultDev
	}

	return rt, nil
}
