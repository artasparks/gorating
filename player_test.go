package gorating

import (
	"testing"
)

func TestRating(t *testing.T) {
	// t.Errorf("Test Failed")
	var r CompareablePlayer = NewBasePlayer("zed")
	if r == nil {
		t.Errorf("Test Failed: nil")
	}
}
