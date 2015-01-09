package gorating

import (
	"testing"
)

func TestRating(t *testing.T) {
	// t.Errorf("Test Failed")
	var r Player = DefaultBasePlayer("zed")
	if r == nil {
		t.Errorf("Test Failed: nil")
	}
}
