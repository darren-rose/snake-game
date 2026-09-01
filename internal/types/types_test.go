package types

import (
	"testing"
	"time"
)

// TestPointEquality tests Point struct equality
func TestPointEquality(t *testing.T) {
	p1 := Point{X: 5, Y: 10}
	p2 := Point{X: 5, Y: 10}
	p3 := Point{X: 5, Y: 11}

	if p1 != p2 {
		t.Errorf("Points with same coordinates should be equal: %v != %v", p1, p2)
	}

	if p1 == p3 {
		t.Errorf("Points with different coordinates should not be equal: %v == %v", p1, p3)
	}
}

// TestExplosionCreation tests Explosion struct initialization
func TestExplosionCreation(t *testing.T) {
	exp := Explosion{
		X:        100.5,
		Y:        200.5,
		Radius:   10.0,
		TimeLeft: 300 * time.Millisecond,
	}

	if exp.X != 100.5 {
		t.Errorf("Expected X=100.5, got %v", exp.X)
	}

	if exp.Radius <= 0 {
		t.Errorf("Expected positive radius, got %v", exp.Radius)
	}
}

// TestGameInitialization tests Game struct initialization
func TestGameInitialization(t *testing.T) {
	g := &Game{
		Snake:          []Point{{X: 0, Y: 0}},
		Lives:          3,
		Score:          0,
		Level:          1,
		GameOver:       false,
		UpdateInterval: 200 * time.Millisecond,
	}

	if g.Lives != 3 {
		t.Errorf("Expected lives=3, got %d", g.Lives)
	}

	if len(g.Snake) == 0 {
		t.Error("Snake should not be empty")
	}

	if g.GameOver {
		t.Error("Game should not be over initially")
	}
}
