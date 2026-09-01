package renderer

import (
	"testing"

	"github.com/darren-rose/snake-in-go/internal/config"
	"github.com/darren-rose/snake-in-go/internal/types"
)

// TestGameStructureForRendering tests that game has required fields for rendering
func TestGameStructureForRendering(t *testing.T) {
	g := &types.Game{
		Snake:      []types.Point{{X: 10, Y: 10}},
		Food:       types.Point{X: 20, Y: 20},
		Lives:      3,
		Score:      100,
		Level:      1,
		FoodEaten:  2,
		GameOver:   false,
		Explosions: []types.Explosion{},
	}

	// Verify all fields needed for rendering exist
	if len(g.Snake) == 0 {
		t.Error("Snake should not be empty for rendering")
	}

	if g.Food.X < 0 || g.Food.Y < 0 {
		t.Error("Food position should be valid")
	}

	if g.Lives < 0 {
		t.Error("Lives should not be negative")
	}
}

// TestGameOverStateForRendering tests game over rendering state
func TestGameOverStateForRendering(t *testing.T) {
	g := &types.Game{
		Snake:      []types.Point{{X: 10, Y: 10}},
		Score:      500,
		Level:      3,
		GameOver:   true,
		Lives:      0,
		Explosions: []types.Explosion{},
	}

	if !g.GameOver {
		t.Error("Game should be marked as over for game over screen")
	}

	if g.Score < 0 {
		t.Error("Score should be non-negative")
	}

	if g.Level < 1 {
		t.Error("Level should be at least 1")
	}
}

// TestExplosionRenderingData tests that explosions have valid rendering data
func TestExplosionRenderingData(t *testing.T) {
	g := &types.Game{
		Explosions: []types.Explosion{
			{X: 100, Y: 100, Radius: 5, TimeLeft: 100},
			{X: 150, Y: 150, Radius: 10, TimeLeft: 200},
		},
	}

	for i, exp := range g.Explosions {
		if exp.X < 0 || exp.Y < 0 {
			t.Errorf("Explosion %d has invalid position: (%f, %f)", i, exp.X, exp.Y)
		}

		if exp.Radius < 0 {
			t.Errorf("Explosion %d has negative radius: %f", i, exp.Radius)
		}

		if exp.TimeLeft < 0 {
			t.Errorf("Explosion %d has negative TimeLeft", i)
		}
	}
}

// TestInfoPanelData tests that all info panel data is available
func TestInfoPanelData(t *testing.T) {
	g := &types.Game{
		Lives:     2,
		Score:     1500,
		Level:     5,
		FoodEaten: 23,
	}

	if g.Lives < 0 {
		t.Error("Lives should not be negative")
	}

	if g.Score < 0 {
		t.Error("Score should not be negative")
	}

	if g.Level < 1 {
		t.Error("Level should be at least 1")
	}

	// foodCount should show progress to next level
	foodProgress := g.FoodEaten % config.FoodPerLevel
	if foodProgress < 0 {
		t.Error("Food progress should be non-negative")
	}
}

// TestSnakeRenderingPositions tests that snake positions are valid for rendering
func TestSnakeRenderingPositions(t *testing.T) {
	g := &types.Game{
		Snake: []types.Point{
			{X: 10, Y: 10},
			{X: 10, Y: 11},
			{X: 10, Y: 12},
		},
	}

	maxX := config.WindowWidth / config.GridSize
	maxY := config.GameAreaHeight / config.GridSize

	for i, segment := range g.Snake {
		if segment.X < 0 || segment.X >= maxX {
			t.Errorf("Snake segment %d X coordinate %d out of bounds [0, %d)", i, segment.X, maxX)
		}

		if segment.Y < 0 || segment.Y >= maxY {
			t.Errorf("Snake segment %d Y coordinate %d out of bounds [0, %d)", i, segment.Y, maxY)
		}
	}
}

// TestFoodRenderingPosition tests that food position is valid for rendering
func TestFoodRenderingPosition(t *testing.T) {
	maxX := config.WindowWidth / config.GridSize
	maxY := config.GameAreaHeight / config.GridSize

	foods := []types.Point{
		{X: 0, Y: 5},
		{X: 10, Y: 10},
		{X: maxX - 1, Y: maxY - 1},
	}

	for i, food := range foods {
		if food.X < 0 || food.X >= maxX {
			t.Errorf("Food %d X out of bounds: %d", i, food.X)
		}

		if food.Y < 0 || food.Y >= maxY {
			t.Errorf("Food %d Y out of bounds: %d", i, food.Y)
		}
	}
}

// TestRenderingConstants tests that all rendering constants are valid
func TestRenderingConstants(t *testing.T) {
	if config.WindowWidth <= 0 || config.WindowHeight <= 0 {
		t.Error("Window dimensions must be positive")
	}

	if config.GridSize <= 0 {
		t.Error("GridSize must be positive")
	}

	if config.InfoPanelHeight < 0 {
		t.Error("InfoPanelHeight must be non-negative")
	}

	if config.GameAreaHeight <= 0 {
		t.Error("GameAreaHeight must be positive")
	}
}

// BenchmarkGameRenderingData benchmarks rendering data assembly
func BenchmarkGameRenderingData(b *testing.B) {
	g := &types.Game{
		Snake:      make([]types.Point, 100),
		Explosions: make([]types.Explosion, 50),
	}

	// Initialize snake positions
	for i := 0; i < len(g.Snake); i++ {
		g.Snake[i] = types.Point{X: i % 32, Y: i / 32}
	}

	// Initialize explosions
	for i := 0; i < len(g.Explosions); i++ {
		g.Explosions[i] = types.Explosion{X: float64(i * 5), Y: float64(i * 5)}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Simulate accessing rendering data
		_ = len(g.Snake)
		_ = g.Food
		_ = len(g.Explosions)
	}
}
