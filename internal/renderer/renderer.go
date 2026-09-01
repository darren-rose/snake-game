// Package renderer handles all game rendering
package renderer

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/darren-rose/snake-in-go/internal/config"
	"github.com/darren-rose/snake-in-go/internal/types"
)

// Draw renders the game to the screen
func Draw(screen *ebiten.Image, g *types.Game) {
	drawInfoPanel(screen, g)
	drawGameArea(screen)
	drawSnake(screen, g)
	drawFood(screen, g)
	drawExplosions(screen, g)

	if g.GameOver {
		drawGameOverScreen(screen, g)
	}
}

// drawInfoPanel renders the top information panel
func drawInfoPanel(screen *ebiten.Image, g *types.Game) {
	// Panel background
	vector.DrawFilledRect(
		screen, 0, 0,
		float32(config.WindowWidth), float32(config.InfoPanelHeight),
		config.ColorInfoPanel, true,
	)

	// Panel text
	panelText := fmt.Sprintf(
		"Lives: %d | Score: %d | Level: %d | Food: %d/%d",
		g.Lives, g.Score, g.Level, g.FoodEaten%config.FoodPerLevel, config.FoodPerLevel,
	)
	ebitenutil.DebugPrintAt(screen, panelText, 5, 5)
}

// drawGameArea renders the game area background
func drawGameArea(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen, 0, float32(config.InfoPanelHeight),
		float32(config.WindowWidth), float32(config.GameAreaHeight),
		config.ColorGameArea, true,
	)
}

// drawSnake renders the snake
func drawSnake(screen *ebiten.Image, g *types.Game) {
	for _, segment := range g.Snake {
		x := float32(segment.X * config.GridSize)
		y := float32(segment.Y*config.GridSize + config.InfoPanelHeight)

		vector.DrawFilledRect(
			screen, x, y,
			float32(config.GridSize), float32(config.GridSize),
			config.ColorSnake, true,
		)
	}
}

// drawFood renders the food
func drawFood(screen *ebiten.Image, g *types.Game) {
	x := float32(g.Food.X * config.GridSize)
	y := float32(g.Food.Y*config.GridSize + config.InfoPanelHeight)

	vector.DrawFilledRect(
		screen, x, y,
		float32(config.GridSize), float32(config.GridSize),
		config.ColorFood, true,
	)
}

// drawExplosions renders all active explosions
func drawExplosions(screen *ebiten.Image, g *types.Game) {
	for _, exp := range g.Explosions {
		// Calculate alpha based on time remaining
		alpha := uint8(float64(exp.TimeLeft) / float64(300*time.Millisecond) * 255)

		vector.DrawFilledCircle(
			screen,
			float32(exp.X),
			float32(exp.Y+float64(config.InfoPanelHeight)),
			float32(exp.Radius),
			color.RGBA{config.ColorExplosion.R, config.ColorExplosion.G, config.ColorExplosion.B, alpha},
			true,
		)
	}
}

// drawGameOverScreen renders the game over screen
func drawGameOverScreen(screen *ebiten.Image, g *types.Game) {
	gameOverText := fmt.Sprintf(
		"GAME OVER! Final Score: %d | Level: %d",
		g.Score, g.Level,
	)
	ebitenutil.DebugPrintAt(screen, gameOverText, config.WindowWidth/2-130, config.WindowHeight/2-10)

	instructionsText := "Press R to restart or Q to quit"
	ebitenutil.DebugPrintAt(screen, instructionsText, config.WindowWidth/2-130, config.WindowHeight/2+10)
}
