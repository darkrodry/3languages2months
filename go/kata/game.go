package kata

type Game struct {
	points int
}

func(g* Game) roll(pins int) {
	g.points += pins
}

func(g* Game) score() int {
	return g.points
}
