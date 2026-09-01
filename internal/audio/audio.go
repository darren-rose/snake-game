// Package audio contains sound effect generation and playback
package audio

import (
	"bytes"
	"math"

	"github.com/hajimehoshi/ebiten/v2/audio"

	"github.com/darren-rose/snake-in-go/internal/config"
	"github.com/darren-rose/snake-in-go/internal/types"
)

// PlayEatSound plays the food eating sound
func PlayEatSound(g *types.Game) {
	playTone(g, config.EatFrequency, config.EatDuration, config.EatVolume)
}

// PlayDieSound plays the collision/death sound
func PlayDieSound(g *types.Game) {
	playTone(g, config.DieFrequency, config.DieDuration, config.DieVolume)
}

// PlayMoveSound plays the snake movement sound
func PlayMoveSound(g *types.Game) {
	playTone(g, config.MoveFrequency, config.MoveDuration, config.MoveVolume)
}

// playTone is a helper that generates and plays a tone
func playTone(g *types.Game, frequency, durationMs int, volume float64) {
	if g.AudioContext == nil {
		return
	}

	samples := generateToneSamples(frequency, durationMs)
	reader := bytes.NewReader(samples)

	player, err := audio.NewPlayer(g.AudioContext, reader)
	if err != nil {
		return
	}

	player.SetVolume(volume)
	player.Play()
}

// generateToneSamples generates audio samples for a sine wave tone
func generateToneSamples(frequency, durationMs int) []byte {
	const BitsPerSample = 16
	const Amplitude = 0.3

	sampleCount := config.AudioSampleRate * durationMs / 1000
	buf := make([]byte, sampleCount*2) // 16-bit = 2 bytes per sample

	for i := 0; i < sampleCount; i++ {
		// Generate sine wave
		sineValue := math.Sin(2 * math.Pi * float64(frequency) * float64(i) / float64(config.AudioSampleRate))
		sample := int16(sineValue * Amplitude * 32767)

		// Convert to little-endian bytes
		buf[i*2] = byte(sample & 0xFF)
		buf[i*2+1] = byte((sample >> 8) & 0xFF)
	}

	return buf
}
