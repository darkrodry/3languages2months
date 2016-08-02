package kata

type Game struct {
	currentRoll int
	rolls [21]int
}

func(g* Game) roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll += 1
}

func(g* Game) score() int {
	score := 0
	for frame := 0; frame < 10; frame++ {
		roll := g.rolls[frame*2] + g.rolls[frame*2+1]
		score += roll
		if roll == 10 {
			score +=  g.rolls[frame*2+2]
		}
	}
	return score
}
