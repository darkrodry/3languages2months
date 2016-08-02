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
		if g.rolls[frame*2] == 10 {
			score +=  g.rolls[frame*2+2] + g.rolls[frame*2+3]
		} else if roll == 10 {
			score += g.rolls[frame*2+2]
		}
	}
	return score
}
