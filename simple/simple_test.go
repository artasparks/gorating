package simple

import (
	"github.com/Kashomon/gorating"
	"testing"
)

func TestPlayerMap(t *testing.T) {
	g := []gorating.Game{
		&Game{
			&Player{"1", 100},
			&Player{"2", 200},
			1.0,
		},
		&Game{
			&Player{"2", 200},
			&Player{"3", 300},
			1.0,
		},
	}
	if g == nil {
		t.Errorf("Test Failed")
	}

	m := gorating.PlayerMap(g)

	if len(m["1"]) != 1 {
		t.Errorf("Test Failed")
	}
	if len(m["2"]) != 2 {
		t.Errorf("Test Failed")
	}
	if len(m["3"]) != 1 {
		t.Errorf("Test Failed")
	}
}
