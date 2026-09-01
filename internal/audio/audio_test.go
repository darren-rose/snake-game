package audio

import (
	"testing"

	"github.com/darren-rose/snake-in-go/internal/config"
	"github.com/darren-rose/snake-in-go/internal/types"
)

// TestGenerateToneSamples tests audio sample generation
func TestGenerateToneSamples(t *testing.T) {
	tests := []struct {
		name       string
		frequency  int
		durationMs int
		minLength  int
		maxLength  int
	}{
		{
			name:       "Move sound",
			frequency:  config.MoveFrequency,
			durationMs: config.MoveDuration,
			minLength:  1,
			maxLength:  10000,
		},
		{
			name:       "Eat sound",
			frequency:  config.EatFrequency,
			durationMs: config.EatDuration,
			minLength:  1,
			maxLength:  50000,
		},
		{
			name:       "Die sound",
			frequency:  config.DieFrequency,
			durationMs: config.DieDuration,
			minLength:  1,
			maxLength:  100000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			samples := generateToneSamples(tt.frequency, tt.durationMs)

			if len(samples) < tt.minLength {
				t.Errorf("Sample length %d is less than minimum %d", len(samples), tt.minLength)
			}

			if len(samples) > tt.maxLength {
				t.Errorf("Sample length %d exceeds maximum %d", len(samples), tt.maxLength)
			}

			// Samples should be 16-bit (2 bytes per sample)
			if len(samples)%2 != 0 {
				t.Error("Sample length should be even (16-bit samples)")
			}
		})
	}
}

// TestGenerateToneFrequencies tests that different frequencies generate samples
func TestGenerateToneFrequencies(t *testing.T) {
	frequencies := []int{100, 500, 800, 1000, 2000}

	for _, freq := range frequencies {
		samples := generateToneSamples(freq, 100)

		if len(samples) == 0 {
			t.Errorf("No samples generated for frequency %d", freq)
		}
	}
}

// TestPlaySoundsWithNilContext tests sound playback with nil audio context
func TestPlaySoundsWithNilContext(t *testing.T) {
	g := &types.Game{
		AudioContext: nil,
	}

	// These should not panic even with nil context
	PlayMoveSound(g)
	PlayEatSound(g)
	PlayDieSound(g)
}

// TestSampleRateCalculation tests audio sample calculations
func TestSampleRateCalculation(t *testing.T) {
	const durationMs = 100
	expectedSamples := config.AudioSampleRate * durationMs / 1000

	samples := generateToneSamples(440, durationMs)
	actualSamples := len(samples) / 2 // 16-bit = 2 bytes per sample

	if actualSamples < expectedSamples-2 || actualSamples > expectedSamples+2 {
		t.Errorf("Expected ~%d samples for %dms at %dHz, got %d",
			expectedSamples, durationMs, config.AudioSampleRate, actualSamples)
	}
}

// BenchmarkGenerateToneSamples benchmarks audio generation
func BenchmarkGenerateToneSamples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateToneSamples(config.EatFrequency, config.EatDuration)
	}
}
