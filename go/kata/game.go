package kata

type Game struct {
	currentRoll int
	rolls [21]int
}

func(g* Game) roll(pins int) {
	g.rolls[g.currentRoll] = pins
	if pins == 10 && g.currentRoll % 2 == 0 {
		g.currentRoll += 2
	} else {
		g.currentRoll += 1
	}
}

func(g* Game) score() int {
	score := 0
	for frame := 0; frame < 10; frame++ {
		roll := g.rolls[frame*2] + g.rolls[frame*2+1]
		score += roll
		if isStrike(g, frame) {
			score +=  g.rolls[frame*2+2] + g.rolls[frame*2+3]
		} else if isSpare(g, frame) {
			score += g.rolls[frame*2+2]
		}
	}
	return score
}

func isStrike(g* Game, frame int)  bool {
	return g.rolls[frame*2] == 10
}

func isSpare(g* Game, frame int) bool {
	return g.rolls[frame*2] + g.rolls[frame*2] == 10 &&
		g.rolls[frame*2] != 10
}
