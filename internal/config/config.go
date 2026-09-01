// Package config contains all game configuration constants
package config

import (
	"image/color"
	"time"

	"github.com/darren-rose/snake-in-go/internal/types"
)

// Window dimensions
const (
	WindowWidth  = 320
	WindowHeight = 240
)

// Game grid
const (
	GridSize        = 5
	InfoPanelHeight = 30
	GameAreaHeight  = WindowHeight - InfoPanelHeight
)

// Game speed
const (
	BaseUpdateInterval = time.Millisecond * 200
	SpeedUpPercent     = 0.1 // 10% speed increase per level
)

// Audio configuration
const (
	AudioSampleRate = 44100
)

// Game mechanics
const (
	InitialLives    = 3
	FoodPerLevel    = 5
	FoodEatPoints   = 20
	SurvivalPoints  = 1
	SurvivalSeconds = 5
)

// Sound parameters
const (
	MoveFrequency = 500
	MoveDuration  = 30
	MoveVolume    = 0.1
	EatFrequency  = 800
	EatDuration   = 100
	EatVolume     = 0.3
	DieFrequency  = 300
	DieDuration   = 200
	DieVolume     = 0.5
)

// Color constants
var (
	ColorSnake     = color.RGBA{0, 255, 0, 255}   // Green
	ColorFood      = color.RGBA{255, 0, 0, 255}   // Red
	ColorExplosion = color.RGBA{255, 165, 0, 255} // Orange
	ColorInfoPanel = color.RGBA{50, 50, 50, 255}  // Dark gray
	ColorGameArea  = color.RGBA{20, 20, 20, 255}  // Darker gray
)

// Direction vectors
var (
	DirectionUp    = types.Point{Y: -1}
	DirectionDown  = types.Point{Y: 1}
	DirectionLeft  = types.Point{X: -1}
	DirectionRight = types.Point{X: 1}
)
