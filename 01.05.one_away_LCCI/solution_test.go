package one_away_LCCI

import (
	"testing"
)

func TestOneEditAway(t *testing.T) {
	if !oneEditAway("pale", "ple") {
		t.Errorf("error")
	}
	if oneEditAway("pales", "pal") {
		t.Error("error")
	}
	if oneEditAway("horse", "ros") {
		t.Error("error")
	}
	if oneEditAway("intention", "execution") {
		t.Error("error")
	}
}
