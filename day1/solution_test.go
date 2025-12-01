package main

import (
	"testing"
)

func TestRightRotationToFloor(t *testing.T) {
	dial := newDial()
	dial.position = 52
	dial.moveDial("R", 48)
	want := 0
	got := dial.position

	if got != want {
		t.Errorf("Expected position at %d, but got %d \n", want, got)
	}

}

func TestLeftRotationToCieling(t *testing.T) {
	dial := newDial()
	dial.position = 55
	dial.moveDial("L", 55)
	want := 0
	got := dial.position

	if got != want {
		t.Errorf("Expected position at %d, but got %d \n", want, got)
	}

}

func TestLeftRotationPastCieling(t *testing.T) {
	dial := newDial()
	dial.moveDial("L", 68)
	want := 82
	got := dial.position

	if got != want {
		t.Errorf("Expected position at %d, but got %d \n", want, got)
	}

}

func TestRightRotationPastFloor(t *testing.T) {
	dial := newDial()
	dial.position = 95
	dial.moveDial("R", 60)
	want := 55
	got := dial.position

	if got != want {
		t.Errorf("Expected position at %d, but got %d \n", want, got)
	}

}

func TestRotationsWithClicks(t *testing.T) {
	wantRotationLandOnZeroCount := 0
	wantFullRotationCount := 0
	dial := newDial()
	dial.moveDial("L", 68)
	wantFullRotationCount += 1
	dial.moveDial("L", 30)
	dial.moveDial("R", 48)
	wantRotationLandOnZeroCount += 1
	dial.moveDial("L", 5)
	dial.moveDial("R", 60)
	wantFullRotationCount += 1
	dial.moveDial("L", 55)
	wantRotationLandOnZeroCount += 1
	dial.moveDial("L", 1)
	dial.moveDial("L", 99)
	wantRotationLandOnZeroCount += 1
	dial.moveDial("R", 14)
	dial.moveDial("L", 82)
	wantFullRotationCount += 1

	if dial.RotationLandOnZeroCount != wantRotationLandOnZeroCount {
		t.Errorf("Expected RotationLandOnZeroCount at %d, but got %d \n", wantRotationLandOnZeroCount, dial.RotationLandOnZeroCount)

	}

	if dial.FullRotationCount != wantFullRotationCount {
		t.Errorf("Expected FullRotationCount at %d, but got %d \n", wantFullRotationCount, dial.FullRotationCount)

	}

	if dial.position != 32 {
		t.Errorf("Expected position at %d, but got %d \n", 32, dial.position)

	}

}
