package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"

	"github.com/darren-rose/snake-in-go/internal/config"
	"github.com/darren-rose/snake-in-go/internal/game"
	"github.com/darren-rose/snake-in-go/internal/renderer"
	"github.com/darren-rose/snake-in-go/internal/types"
)

// Application wraps the game and implements ebiten.Game interface
type Application struct {
	game *types.Game
}

// Update updates the game state
func (app *Application) Update() error {
	return game.Update(app.game)
}

// Draw renders the game
func (app *Application) Draw(screen *ebiten.Image) {
	renderer.Draw(screen, app.game)
}

// Layout returns the game window size
func (app *Application) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.WindowWidth, config.WindowHeight
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Initialize audio
	audioCtx := audio.NewContext(config.AudioSampleRate)

	// Create game
	g := game.NewGame(audioCtx)

	// Wrap game in application
	app := &Application{
		game: g,
	}

	// Configure window
	ebiten.SetWindowSize(config.WindowWidth, config.WindowHeight)
	ebiten.SetWindowTitle("Snake Game")

	// Run game
	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
