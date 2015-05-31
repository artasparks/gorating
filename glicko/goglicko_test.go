package goglicko

import "testing"

// Much of this data comes from the paper:
// http://en.wikipedia.org/wiki/Glicko_rating_system
var pl = NewRating(1500, 200, DefaultVol)
var opps = []*Rating{
	NewRating(1400, 30, DefaultVol),
	NewRating(1550, 100, DefaultVol),
	NewRating(1700, 300, DefaultVol),
}
var results = []Result{1, 0, 0}

func TestEquivTransfOpps(t *testing.T) {
	for i := range opps {
		o := opps[i]
		o2 := opps[i].ToGlicko2().FromGlicko2()
		if !o.MostlyEquals(o2, 0.0001) {
			t.Errorf("o %v != o2 %v", o, o2)
		}
	}
}

func TestToGlicko2(t *testing.T) {
	p2 := pl.ToGlicko2()
	exp := NewRating(0, 1.1513, DefaultVol)
	if !p2.MostlyEquals(exp, 0.0001) {
		t.Errorf("p2 %v != expected %v", p2, exp)
	}
}

func TestOppToGlicko2(t *testing.T) {
	exp := []*Rating{
		NewRating(-0.5756, 0.1727, DefaultVol),
		NewRating(0.2878, 0.5756, DefaultVol),
		NewRating(1.1513, 1.7269, DefaultVol),
	}
	for i := range exp {
		g2 := opps[i].ToGlicko2()
		if !g2.MostlyEquals(exp[i], 0.0001) {
			t.Errorf("For i=%v: Glicko2 scaled opp %v != expected %v\n", i, g2, exp[i])
		}
	}
}

func TestEeGeeValues(t *testing.T) {
	expGee := []float64{0.9955, 0.9531, 0.7242}
	expEe := []float64{0.639, 0.432, 0.303}
	p2 := pl.ToGlicko2()
	for i := range opps {
		o := opps[i].ToGlicko2()
		geeVal := gee(o.Deviation)
		if !floatsMostlyEqual(geeVal, expGee[i], 0.0001) {
			t.Errorf("Floats not mostly equal. g=%v exp_g=%v", geeVal, expGee[i])
		}
		eeVal := ee(p2.Rating, o.Rating, o.Deviation)
		if !floatsMostlyEqual(eeVal, expEe[i], 0.001) {
			t.Errorf("Floats not mostly equal. ee=%v exp_ee=%v", eeVal, expEe[i])
		}
	}
}

func TestAlgorithm(t *testing.T) {
	p2 := pl.ToGlicko2()
	gees := make([]float64, len(opps))
	ees := make([]float64, len(opps))
	for i := range opps {
		o := opps[i].ToGlicko2()
		gees[i] = gee(o.Deviation)
		ees[i] = ee(p2.Rating, o.Rating, o.Deviation)
	}
	estVar := estVariance(gees, ees)
	exp := 1.7785
	if !floatsMostlyEqual(estVar, exp, 0.001) {
		t.Errorf("estvar %v != exp %v", estVar, exp)
	}

	// Test Delta
	estImpPart := estImprovePartial(gees, ees, results)
	estImp := estVar * estImpPart
	expEstImp := -0.4834
	if !floatsMostlyEqual(estImp, expEstImp, 0.001) {
		t.Errorf("delta %v != exp %v", estImp, expEstImp)
	}

	// Test calculating the new volatility
	newVol := newVolatility(estVar, estImp, p2)
	expNewVol := 0.05999
	if !floatsMostlyEqual(newVol, expNewVol, 0.0001) {
		t.Errorf("newVol %v != expNewVol %v", newVol, expNewVol)
	}

	newDev := newDeviation(p2.Deviation, newVol, estVar)
	expNewDev := 0.8722
	if !floatsMostlyEqual(newDev, expNewDev, 0.0001) {
		t.Errorf("newDev %v != expNewDev %v", newDev, expNewDev)
	}

	newRating := newRatingVal(p2.Rating, newDev, estImpPart)
	expNewRating := -0.2069
	if !floatsMostlyEqual(newRating, expNewRating, 0.0001) {
		t.Errorf("newRating %v != expNewRating %v", newRating, expNewRating)
	}

	newPlayer := NewRating(newRating, newDev, newVol).FromGlicko2()
	expNewRatingV1 := 1464.06
	if !floatsMostlyEqual(newPlayer.Rating, expNewRatingV1, 0.01) {
		t.Errorf("newPlayer.Rating %v != expNewRatingV1 %v",
			newPlayer.Rating, expNewRatingV1)
	}
	expNewDevV1 := 151.52
	if !floatsMostlyEqual(newPlayer.Deviation, expNewDevV1, 0.01) {
		t.Errorf("newPlayer.Deviation %v != expNewDevV1 %v",
			newPlayer.Deviation, expNewDevV1)
	}
}

func TestCalculateRating(t *testing.T) {
	r, err := CalculateRating(pl, opps, results)
	if err != nil {
		t.Errorf("Error while calculating results: %v", err)
		return
	}

	expNewVol := 0.05999
	if !floatsMostlyEqual(r.Volatility, expNewVol, 0.0001) {
		t.Errorf("r.Volatility %v != expNewVol %v", r.Volatility, expNewVol)
	}
	expNewRatingV1 := 1464.06
	if !floatsMostlyEqual(r.Rating, expNewRatingV1, 0.01) {
		t.Errorf("r.Rating %v != expNewRatingV1 %v",
			r.Rating, expNewRatingV1)
	}
	expNewDevV1 := 151.52
	if !floatsMostlyEqual(r.Deviation, expNewDevV1, 0.01) {
		t.Errorf("r.Deviation %v != expNewDevV1 %v",
			r.Deviation, expNewDevV1)
	}
}
