package goglicko

import "testing"

func BenchmarkSimpleExample(b *testing.B) {
	p := DefaultRating()
	o := []*Rating{
		NewRating(1400, 30, DefaultVol),
		NewRating(1550, 100, DefaultVol),
		NewRating(1700, 300, DefaultVol),
	}
	res := []Result{1, 0, 0}

	for i := 0; i < b.N; i++ {
		CalculateRating(p, o, res)
	}
}
