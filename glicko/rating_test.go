package goglicko

import "testing"

func TestScaleRescale(t *testing.T) {
	def := DefaultRating()
	p2 := def.ToGlicko2().FromGlicko2()
	if !def.MostlyEquals(p2, 0.0001) {
		t.Errorf("Test Failed. def %v != p2 %v", def, p2)
	}
}

func TestStringyf(t *testing.T) {
	def := DefaultRating()
	if def.String() != "{Rating[1500.000] Deviation[350.000] Volatility[0.060]}" {
		t.Errorf("Error. String form was %v", def.String())
	}
}
