package game

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var g Game

func rollNPins(frame int, pins int) {
	for i := 0; i < frame; i++ {
		g.Roll(pins)
	}
}

func TestMain(m *testing.M) {
	g = Game{}
	os.Exit(m.Run())
}

func TestRollAllZero(t *testing.T) {
	rollNPins(20, 0)
	assert.Equal(t, 0, g.Score())
}

func TestRollAllOne(t *testing.T) {
	rollNPins(20, 1)
	assert.Equal(t, 20, g.Score())
}
