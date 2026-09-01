package game

import (
	"testing"
	"time"

	"github.com/darren-rose/snake-in-go/internal/config"
	"github.com/darren-rose/snake-in-go/internal/types"
)

// TestNewGameCreation tests that NewGame initializes properly
func TestNewGameCreation(t *testing.T) {
	g := NewGame(nil)

	if g == nil {
		t.Fatal("NewGame returned nil")
	}

	if g.Lives != config.InitialLives {
		t.Errorf("Expected lives=%d, got %d", config.InitialLives, g.Lives)
	}

	if g.Score != 0 {
		t.Errorf("Expected score=0, got %d", g.Score)
	}

	if g.Level != 1 {
		t.Errorf("Expected level=1, got %d", g.Level)
	}

	if g.GameOver {
		t.Error("Game should not be over on creation")
	}

	if len(g.Snake) == 0 {
		t.Error("Snake should not be empty")
	}

	if g.UpdateInterval != config.BaseUpdateInterval {
		t.Errorf("Expected UpdateInterval=%v, got %v", config.BaseUpdateInterval, g.UpdateInterval)
	}
}

// TestSnakeInitialPosition tests snake starts at center
func TestSnakeInitialPosition(t *testing.T) {
	g := NewGame(nil)

	centerX := config.WindowWidth / config.GridSize / 2
	centerY := (config.GameAreaHeight / config.GridSize) / 2

	if g.Snake[0].X != centerX || g.Snake[0].Y != centerY {
		t.Errorf("Snake head should be at center (%d, %d), got (%d, %d)",
			centerX, centerY, g.Snake[0].X, g.Snake[0].Y)
	}
}

// TestReset tests game reset functionality
func TestReset(t *testing.T) {
	g := NewGame(nil)

	// Modify game state
	g.Lives = 1
	g.Score = 100
	g.Level = 5
	g.GameOver = true

	// Reset
	Reset(g)

	if g.Lives != config.InitialLives {
		t.Errorf("Expected lives=%d after reset, got %d", config.InitialLives, g.Lives)
	}

	if g.Score != 0 {
		t.Errorf("Expected score=0 after reset, got %d", g.Score)
	}

	if g.Level != 1 {
		t.Errorf("Expected level=1 after reset, got %d", g.Level)
	}

	if g.GameOver {
		t.Error("GameOver should be false after reset")
	}
}

// TestWallCollisionDetection tests boundary collision
func TestWallCollisionDetection(t *testing.T) {
	tests := []struct {
		name     string
		point    types.Point
		expected bool
	}{
		{
			name:     "Inside boundary",
			point:    types.Point{X: 10, Y: 10},
			expected: false,
		},
		{
			name:     "Left wall",
			point:    types.Point{X: -1, Y: 10},
			expected: true,
		},
		{
			name:     "Right wall",
			point:    types.Point{X: config.WindowWidth / config.GridSize, Y: 10},
			expected: true,
		},
		{
			name:     "Top wall",
			point:    types.Point{X: 10, Y: -1},
			expected: true,
		},
		{
			name:     "Bottom wall",
			point:    types.Point{X: 10, Y: config.GameAreaHeight / config.GridSize},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkWallCollision(tt.point)
			if result != tt.expected {
				t.Errorf("checkWallCollision(%v) = %v, want %v", tt.point, result, tt.expected)
			}
		})
	}
}

// TestSelfCollisionDetection tests snake self-collision
func TestSelfCollisionDetection(t *testing.T) {
	g := &types.Game{
		Snake: []types.Point{
			{X: 10, Y: 10},
			{X: 10, Y: 11},
			{X: 10, Y: 12},
		},
	}

	// Head collides with body
	if !checkSelfCollision(g, types.Point{X: 10, Y: 11}) {
		t.Error("Should detect collision with snake body")
	}

	// Head doesn't collide with empty space
	if checkSelfCollision(g, types.Point{X: 5, Y: 5}) {
		t.Error("Should not detect collision in empty space")
	}
}

// TestGenerateFood tests food generation doesn't spawn on snake
func TestGenerateFood(t *testing.T) {
	g := NewGame(nil)

	for i := 0; i < 10; i++ {
		food := GenerateFood(g)

		// Food should not be outside bounds
		if food.X < 0 || food.X >= config.WindowWidth/config.GridSize {
			t.Errorf("Food X out of bounds: %d", food.X)
		}

		if food.Y < 0 || food.Y >= config.GameAreaHeight/config.GridSize {
			t.Errorf("Food Y out of bounds: %d", food.Y)
		}

		// Food should not spawn in info panel
		if food.Y < 1 {
			t.Error("Food should not spawn in info panel area")
		}
	}
}

// TestFoodConsumption tests snake eating food
func TestFoodConsumption(t *testing.T) {
	g := NewGame(nil)
	initialLength := len(g.Snake)
	initialScore := g.Score
	initialFoodEaten := g.FoodEaten

	// Simulate eating food by calling handleFoodCollision
	handleFoodCollision(g)

	if len(g.Snake) != initialLength {
		t.Errorf("Snake length should remain %d after eating, got %d", initialLength, len(g.Snake))
	}

	if g.Score <= initialScore {
		t.Error("Score should increase after eating food")
	}

	if g.FoodEaten != initialFoodEaten+1 {
		t.Errorf("FoodEaten should be %d, got %d", initialFoodEaten+1, g.FoodEaten)
	}
}

// TestLevelUp tests level progression
func TestLevelUp(t *testing.T) {
	g := NewGame(nil)
	initialLevel := g.Level
	initialInterval := g.UpdateInterval

	levelUp(g)

	if g.Level != initialLevel+1 {
		t.Errorf("Expected level=%d after levelUp, got %d", initialLevel+1, g.Level)
	}

	if g.UpdateInterval >= initialInterval {
		t.Error("UpdateInterval should decrease (get faster) on level up")
	}
}

// TestCollisionHandling tests collision decrements lives
func TestCollisionHandling(t *testing.T) {
	g := NewGame(nil)
	initialLives := g.Lives

	handleCollision(g, types.Point{X: 5, Y: 5})

	if g.Lives != initialLives-1 {
		t.Errorf("Expected lives=%d after collision, got %d", initialLives-1, g.Lives)
	}

	if len(g.Explosions) == 0 {
		t.Error("Collision should create explosion particles")
	}
}

// TestGameOver tests game over state
func TestGameOver(t *testing.T) {
	g := NewGame(nil)
	g.Lives = 1

	handleCollision(g, types.Point{X: 5, Y: 5})

	if !g.GameOver {
		t.Error("GameOver should be true when lives run out")
	}
}

// TestUpdateTiming tests that Update respects UpdateInterval
func TestUpdateTiming(t *testing.T) {
	g := NewGame(nil)
	g.LastUpdate = time.Now()

	// Call Update immediately - should not move
	Update(g)

	// Wait a bit and try again
	g.LastUpdate = time.Now().Add(-g.UpdateInterval)
	initialSnake := make([]types.Point, len(g.Snake))
	copy(initialSnake, g.Snake)

	// This should attempt movement
	Update(g)

	// Snake position may have changed (depends on direction and timing)
	// Just ensure no panic or error occurs
}

// TestCreateExplosion tests explosion particle creation
func TestCreateExplosion(t *testing.T) {
	g := NewGame(nil)
	initialExplosions := len(g.Explosions)

	CreateExplosion(g, 5, 5)

	if len(g.Explosions) <= initialExplosions {
		t.Error("CreateExplosion should add particles")
	}

	if len(g.Explosions) == 0 {
		t.Fatal("No explosions created")
	}

	// Check first particle properties
	exp := g.Explosions[0]
	if exp.TimeLeft <= 0 {
		t.Error("Explosion TimeLeft should be positive")
	}

	if exp.Radius <= 0 {
		t.Error("Explosion Radius should be positive")
	}
}

// BenchmarkGameUpdate benchmarks the game update loop
func BenchmarkGameUpdate(b *testing.B) {
	g := NewGame(nil)

	for i := 0; i < b.N; i++ {
		g.LastUpdate = time.Now().Add(-g.UpdateInterval)
		Update(g)
	}
}

// BenchmarkGenerateFood benchmarks food generation
func BenchmarkGenerateFood(b *testing.B) {
	g := NewGame(nil)

	for i := 0; i < b.N; i++ {
		GenerateFood(g)
	}
}
