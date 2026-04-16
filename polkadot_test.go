package main

import "testing"

func TestComputePolkadotScore(t *testing.T) {
	if got := computePolkadotScore(); got != 71 {
		t.Errorf("computePolkadotScore() = %d, want 71", got)
	}
}

const minimalArt = "()()\n  ~~~~~~\n O O O\n"

func TestScoreArt(t *testing.T) {
	if got := scoreArt(minimalArt); got != 9 {
		t.Errorf("scoreArt(minimalArt) = %d, want 9", got)
	}
}

func TestLipsRange(t *testing.T) {
	start, end := lipsRange(minimalArt)
	if start != 2 || end != 7 {
		t.Errorf("lipsRange() = (%d, %d), want (2, 7)", start, end)
	}
}

func TestPupilCount(t *testing.T) {
	if got := pupilCount(minimalArt); got != 4 {
		t.Errorf("pupilCount() = %d, want 4", got)
	}
}

func TestCountDots(t *testing.T) {
	inside, outside := countDots(minimalArt, 2, 7)
	if inside != 2 || outside != 1 {
		t.Errorf("countDots(2,7) = inside:%d outside:%d, want inside:2 outside:1",
			inside, outside)
	}
}
