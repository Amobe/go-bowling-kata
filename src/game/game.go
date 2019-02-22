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
		if g.rolls[frameIndex]+g.rolls[frameIndex+1] == 10 {
			g.score += 10 + g.rolls[frameIndex+2]
			frameIndex += 2
		} else {
			g.score += g.rolls[frameIndex]
			frameIndex++
		}

	}
	return g.score
}
