package game

const (
	frameNumberPerGame int = 10
	cleanPinNumber     int = 10
)

// Game represents one round of the bowling game.
type Game struct {
	score   int
	rolls   [21]int
	rollCnt int
}

// Roll counts score with the hitting pins of each roll.
func (g *Game) Roll(pins int) {
	g.rolls[g.rollCnt] = pins
	g.rollCnt++
}

// Score returns the score of the bowling game.
func (g *Game) Score() int {
	rollIndex := 0
	for frame := 0; frame < frameNumberPerGame; frame++ {
		if g.isStrike(rollIndex) {
			g.score += cleanPinNumber + g.getStrikeBonus(rollIndex)
			rollIndex++
		} else if g.isSpare(rollIndex) {
			g.score += cleanPinNumber + g.getSpareBonus(rollIndex)
			rollIndex += 2
		} else {
			g.score += g.getFrameScore(rollIndex)
			rollIndex += 2
		}

	}
	return g.score
}

// isSpare returns true if player cleans 10 pins in 2 rounds.
func (g *Game) isSpare(rollIndex int) bool {
	return g.rolls[rollIndex]+g.rolls[rollIndex+1] == cleanPinNumber
}

// getSpareBonus returns the spare bonus value.
func (g *Game) getSpareBonus(rollIndex int) int {
	return g.rolls[rollIndex+2]
}

// isStrike returns true if player cleans 10 pins in a single round.
func (g *Game) isStrike(rollIndex int) bool {
	return g.rolls[rollIndex] == cleanPinNumber
}

// getStrikeBonus returns the strike bonus value.
func (g *Game) getStrikeBonus(rollIndex int) int {
	return g.rolls[rollIndex+1] + g.rolls[rollIndex+2]
}

// getFrameScore returns the score of the frame without any bonus.
func (g *Game) getFrameScore(rollIndex int) int {
	return g.rolls[rollIndex] + g.rolls[rollIndex+1]
}
