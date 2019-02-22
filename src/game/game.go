package game

// Game represents one round of the bowling game.
type Game struct {
	score int
}

// Roll counts score with the hitting pins of each roll.
func (g *Game) Roll(pins int) {
	g.score += pins
}

// Score returns the score of the bowling game.
func (g *Game) Score() int {
	return g.score
}
