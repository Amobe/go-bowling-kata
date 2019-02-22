package game

// Game represents one round of the bowling game.
type Game struct {
	score    int
	rolls    [21]int
	frameCnt int
}

// Roll counts score with the hitting pins of each roll.
func (g *Game) Roll(pins int) {
	g.rolls[g.frameCnt] = pins
	g.frameCnt++
}

// Score returns the score of the bowling game.
func (g *Game) Score() int {
	frameIndex := 0
	for frameIndex < g.frameCnt {
		if g.rolls[frameIndex] == 10 {
			g.score += 10 + g.rolls[frameIndex+1] + g.rolls[frameIndex+2]
			frameIndex++
		} else if g.isSpare(frameIndex) {
			g.score += 10 + g.getSpareBonus(frameIndex)
			frameIndex += 2
		} else {
			g.score += g.rolls[frameIndex]
			frameIndex++
		}

	}
	return g.score
}

// isSpare returns true if player cleans 10 pins in 2 rounds.
func (g *Game) isSpare(frameIndex int) bool {
	return g.rolls[frameIndex]+g.rolls[frameIndex+1] == 10
}

// getSpareBonus returns the spare bonus value.
func (g *Game) getSpareBonus(frameIndex int) int {
	return g.rolls[frameIndex+2]
}
