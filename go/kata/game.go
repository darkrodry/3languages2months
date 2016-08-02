package kata

type Game struct {
	currentRoll int
	rolls [21]int
}

func(g* Game) roll(pins int) {
	g.currentRoll += 1
	g.rolls[g.currentRoll] = pins
}

func(g* Game) score() int {
	score := 0
	for _, pins := range g.rolls {
		score += pins
	}
	return score
}
