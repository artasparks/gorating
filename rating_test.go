package gorating

import (
	"testing"
)

func TestRatingInterface(t *testing.T) {
	var r Rating = DefaultBaseRating()
	if rv := r.Rating(); rv != DefaultRat {
		t.Errorf("Unexpected rating value")
	}
}
