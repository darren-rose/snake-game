// Package game contains core game logic and mechanics
package game

import (
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"

	gameaudio "github.com/darren-rose/snake-in-go/internal/audio"
	"github.com/darren-rose/snake-in-go/internal/config"
	"github.com/darren-rose/snake-in-go/internal/types"
)

// NewGame creates and initializes a new game
func NewGame(audioCtx *audio.Context) *types.Game {
	g := &types.Game{
		AudioContext:   audioCtx,
		Lives:          config.InitialLives,
		Score:          0,
		Level:          1,
		GameOver:       false,
		LastUpdate:     time.Now(),
		LastScoreTime:  time.Now(),
		UpdateInterval: config.BaseUpdateInterval,
		Direction:      config.DirectionRight,
		NextDirection:  config.DirectionRight,
	}

	// Initialize snake at center
	centerX := config.WindowWidth / config.GridSize / 2
	centerY := (config.GameAreaHeight / config.GridSize) / 2
	g.Snake = []types.Point{{X: centerX, Y: centerY}}

	// Spawn initial food
	g.Food = GenerateFood(g)

	return g
}

// Reset resets the game to initial state
func Reset(g *types.Game) {
	centerX := config.WindowWidth / config.GridSize / 2
	centerY := (config.GameAreaHeight / config.GridSize) / 2

	g.Snake = []types.Point{{X: centerX, Y: centerY}}
	g.Direction = config.DirectionRight
	g.NextDirection = config.DirectionRight
	g.Food = GenerateFood(g)
	g.Lives = config.InitialLives
	g.Score = 0
	g.GameOver = false
	g.Explosions = []types.Explosion{}
	g.LastUpdate = time.Now()
	g.LastScoreTime = time.Now()
	g.Level = 1
	g.FoodEaten = 0
	g.UpdateInterval = config.BaseUpdateInterval
}

// Update is called every frame
func Update(g *types.Game) error {
	if g.GameOver {
		return handleGameOverInput(g)
	}

	updateExplosions(g)
	updateScore(g)
	updateSnakeMovement(g)
	handleInput(g)

	return nil
}

// handleGameOverInput processes input when game is over
func handleGameOverInput(g *types.Game) error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		Reset(g)
	}
	return nil
}

// updateExplosions updates all active explosions
func updateExplosions(g *types.Game) {
	for i := 0; i < len(g.Explosions); i++ {
		g.Explosions[i].TimeLeft -= time.Millisecond * 16 // ~60fps
		g.Explosions[i].Radius += 2

		if g.Explosions[i].TimeLeft <= 0 {
			g.Explosions = append(g.Explosions[:i], g.Explosions[i+1:]...)
			i--
		}
	}
}

// updateScore increments score based on survival time
func updateScore(g *types.Game) {
	if time.Since(g.LastScoreTime) >= time.Duration(config.SurvivalSeconds)*time.Second {
		g.LastScoreTime = time.Now()
		g.Score += config.SurvivalPoints
	}
}

// updateSnakeMovement handles all snake movement and collision detection
func updateSnakeMovement(g *types.Game) {
	if time.Since(g.LastUpdate) < g.UpdateInterval {
		return
	}

	g.LastUpdate = time.Now()
	g.Direction = g.NextDirection

	// Calculate new head position
	head := g.Snake[0]
	head.X += g.Direction.X
	head.Y += g.Direction.Y

	// Check collisions
	if checkWallCollision(head) || checkSelfCollision(g, head) {
		handleCollision(g, head)
		return
	}

	// Move snake
	g.Snake = append([]types.Point{head}, g.Snake...)
	gameaudio.PlayMoveSound(g)

	// Handle food collision
	if head == g.Food {
		handleFoodCollision(g)
	} else {
		// Remove tail when not eating
		g.Snake = g.Snake[:len(g.Snake)-1]
	}
}

// checkWallCollision checks if head hit the wall
func checkWallCollision(head types.Point) bool {
	return head.X < 0 ||
		head.X >= config.WindowWidth/config.GridSize ||
		head.Y < 0 ||
		head.Y >= config.GameAreaHeight/config.GridSize
}

// checkSelfCollision checks if head hit the body
func checkSelfCollision(g *types.Game, head types.Point) bool {
	for i := 1; i < len(g.Snake); i++ {
		if g.Snake[i] == head {
			return true
		}
	}
	return false
}

// handleCollision processes collision with wall or self
func handleCollision(g *types.Game, pos types.Point) {
	g.Lives--
	g.DeathPos = pos
	g.DeathTime = time.Now()
	CreateExplosion(g, pos.X, pos.Y)
	gameaudio.PlayDieSound(g)

	if g.Lives <= 0 {
		g.GameOver = true
	} else {
		// Reset snake position
		centerX := config.WindowWidth / config.GridSize / 2
		centerY := (config.GameAreaHeight / config.GridSize) / 2
		g.Snake = []types.Point{{X: centerX, Y: centerY}}
		g.Direction = config.DirectionRight
		g.NextDirection = config.DirectionRight
	}
}

// handleFoodCollision processes eating food
func handleFoodCollision(g *types.Game) {
	g.Score += config.FoodEatPoints
	g.FoodEaten++
	gameaudio.PlayEatSound(g)

	// Level up after eating enough food
	if g.FoodEaten%config.FoodPerLevel == 0 {
		levelUp(g)
	}

	g.Food = GenerateFood(g)
}

// levelUp increases difficulty and level
func levelUp(g *types.Game) {
	g.Level++
	// Increase speed by SpeedUpPercent each level
	g.UpdateInterval = time.Duration(
		float64(config.BaseUpdateInterval) / (1.0 + float64(g.Level-1)*config.SpeedUpPercent),
	)
}

// handleInput processes keyboard input
func handleInput(g *types.Game) {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.NextDirection = config.DirectionUp
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.NextDirection = config.DirectionDown
	}
	if ebiten.IsKeyPressed(ebiten.KeyO) {
		g.NextDirection = config.DirectionLeft
	}
	if ebiten.IsKeyPressed(ebiten.KeyP) {
		g.NextDirection = config.DirectionRight
	}
}

// GenerateFood creates a new food at a random position
func GenerateFood(g *types.Game) types.Point {
	for {
		food := types.Point{
			X: rand.Intn(config.WindowWidth / config.GridSize),
			Y: rand.Intn(config.GameAreaHeight/config.GridSize) + 1, // +1 to avoid info panel
		}
		if !isSnakePosition(g, food) {
			return food
		}
	}
}

// isSnakePosition checks if a point is occupied by the snake
func isSnakePosition(g *types.Game, p types.Point) bool {
	for _, segment := range g.Snake {
		if segment == p {
			return true
		}
	}
	return false
}

// CreateExplosion creates explosion particles at a position
func CreateExplosion(g *types.Game, x, y int) {
	const ExplosionCount = 8
	const ExplosionSize = 5
	const ExplosionDuration = 300 * time.Millisecond

	for i := 0; i < ExplosionCount; i++ {
		exp := types.Explosion{
			X:        float64(x*config.GridSize + config.GridSize/2),
			Y:        float64(y*config.GridSize + config.GridSize/2),
			Radius:   ExplosionSize,
			TimeLeft: ExplosionDuration,
		}
		g.Explosions = append(g.Explosions, exp)
	}
}
