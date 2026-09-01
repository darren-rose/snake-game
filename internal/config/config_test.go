package config

import (
	"image/color"
	"testing"
	"time"
)

// TestWindowDimensions tests window configuration
func TestWindowDimensions(t *testing.T) {
	if WindowWidth <= 0 || WindowHeight <= 0 {
		t.Errorf("Window dimensions must be positive: %dx%d", WindowWidth, WindowHeight)
	}

	if WindowHeight < InfoPanelHeight {
		t.Errorf("Window height %d must be greater than info panel height %d", WindowHeight, InfoPanelHeight)
	}
}

// TestGameAreaHeight tests game area calculation
func TestGameAreaHeight(t *testing.T) {
	expected := WindowHeight - InfoPanelHeight
	if GameAreaHeight != expected {
		t.Errorf("Expected GameAreaHeight=%d, got %d", expected, GameAreaHeight)
	}
}

// TestGridSize tests grid configuration
func TestGridSize(t *testing.T) {
	if GridSize <= 0 {
		t.Errorf("GridSize must be positive, got %d", GridSize)
	}

	if WindowWidth%GridSize != 0 {
		t.Logf("Warning: WindowWidth %d not evenly divisible by GridSize %d", WindowWidth, GridSize)
	}
}

// TestGameMechanics tests game mechanics constants
func TestGameMechanics(t *testing.T) {
	if InitialLives <= 0 {
		t.Errorf("InitialLives must be positive, got %d", InitialLives)
	}

	if FoodPerLevel <= 0 {
		t.Errorf("FoodPerLevel must be positive, got %d", FoodPerLevel)
	}

	if FoodEatPoints <= 0 {
		t.Errorf("FoodEatPoints must be positive, got %d", FoodEatPoints)
	}
}

// TestAudioConfiguration tests audio setup
func TestAudioConfiguration(t *testing.T) {
	if AudioSampleRate <= 0 {
		t.Errorf("AudioSampleRate must be positive, got %d", AudioSampleRate)
	}

	if AudioSampleRate != 44100 {
		t.Logf("Standard sample rate is 44100Hz, got %d", AudioSampleRate)
	}
}

// TestColorConfiguration tests color values
func TestColorConfiguration(t *testing.T) {
	colors := map[string]color.RGBA{
		"ColorSnake":     ColorSnake,
		"ColorFood":      ColorFood,
		"ColorExplosion": ColorExplosion,
		"ColorInfoPanel": ColorInfoPanel,
		"ColorGameArea":  ColorGameArea,
	}

	for name, clr := range colors {
		// Alpha channel should not be fully transparent
		if clr.A == 0 {
			t.Errorf("%s has zero alpha (fully transparent)", name)
		}
	}
}

// TestSpeedupProgression tests speed increase calculation
func TestSpeedupProgression(t *testing.T) {
	if SpeedUpPercent <= 0 || SpeedUpPercent >= 1 {
		t.Errorf("SpeedUpPercent should be between 0 and 1, got %v", SpeedUpPercent)
	}

	// Verify level 2 is faster than level 1
	level1Interval := BaseUpdateInterval
	level2Duration := float64(BaseUpdateInterval) / (1.0 + SpeedUpPercent)
	level2Interval := time.Duration(int64(level2Duration))

	if level2Interval >= level1Interval {
		t.Error("Level 2 should have shorter update interval than level 1")
	}
}

// TestDirections tests direction vectors
func TestDirections(t *testing.T) {
	if DirectionUp.Y >= 0 {
		t.Error("Up direction should have negative Y")
	}

	if DirectionDown.Y <= 0 {
		t.Error("Down direction should have positive Y")
	}

	if DirectionLeft.X >= 0 {
		t.Error("Left direction should have negative X")
	}

	if DirectionRight.X <= 0 {
		t.Error("Right direction should have positive X")
	}
}
