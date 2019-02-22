package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GameTestSuite struct {
	suite.Suite
	g Game
}

func (s *GameTestSuite) SetupTest() {
	s.g = Game{}
}

// rollMany executes several times roll with giving frame value
// each roll hits few pins with giving pins value.
func (s *GameTestSuite) rollMany(frame int, pins int) {
	for i := 0; i < frame; i++ {
		s.g.Roll(pins)
	}
}

// roolSpare executes 2 rolls to finish a spare.
func (s *GameTestSuite) rollSpare() {
	s.g.Roll(2)
	s.g.Roll(8)
}

func (s *GameTestSuite) TestRollAllZero() {
	s.rollMany(20, 0)
	assert.Equal(s.T(), 0, s.g.Score())
}

func (s *GameTestSuite) TestRollAllOne() {
	s.rollMany(20, 1)
	assert.Equal(s.T(), 20, s.g.Score())
}

func (s *GameTestSuite) TestRollSpare() {
	s.rollSpare()
	s.g.Roll(3)
	s.rollMany(17, 0)
	assert.Equal(s.T(), 16, s.g.Score())
}

func (s *GameTestSuite) TestRollStrike() {
	s.g.Roll(10)
	s.g.Roll(3)
	s.g.Roll(5)
	s.rollMany(16, 0)
	assert.Equal(s.T(), 26, s.g.Score())
}

func TestGameTestSuite(t *testing.T) {
	suite.Run(t, new(GameTestSuite))
}
