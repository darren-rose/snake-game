// Package types defines core game types and models
package types

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

// Point represents a coordinate in the game world
type Point struct {
	X, Y int
}

// Explosion represents a visual explosion effect
type Explosion struct {
	X        float64
	Y        float64
	Radius   float64
	TimeLeft time.Duration
}

// Game is the main game state and controller
type Game struct {
	// Snake and movement
	Snake         []Point
	Direction     Point
	NextDirection Point // Buffered direction input

	// Game state
	Lives     int
	Score     int
	Level     int
	FoodEaten int
	GameOver  bool

	// Timing
	LastUpdate     time.Time
	LastScoreTime  time.Time
	UpdateInterval time.Duration

	// Game objects
	Food       Point
	Explosions []Explosion

	// Audio
	AudioContext *audio.Context

	// Collision tracking
	DeathPos  Point
	DeathTime time.Time
}
