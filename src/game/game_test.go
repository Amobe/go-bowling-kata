package game

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var g Game

// rollMany executes several times roll with giving frame value
// each roll hits few pins with giving pins value.
func rollMany(frame int, pins int) {
	for i := 0; i < frame; i++ {
		g.Roll(pins)
	}
}

func TestMain(m *testing.M) {
	g = Game{}
	os.Exit(m.Run())
}

func TestRollAllZero(t *testing.T) {
	rollMany(20, 0)
	assert.Equal(t, 0, g.Score())
}

func TestRollAllOne(t *testing.T) {
	rollMany(20, 1)
	assert.Equal(t, 20, g.Score())
}

func TestRollSpare(t *testing.T) {
	g.Roll(2)
	g.Roll(8)
	g.Roll(3)
	rollMany(17, 0)
	assert.Equal(t, 16, g.Score())
}
