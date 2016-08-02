package kata

import "testing"

var game Game

func setUp() {
	game = Game{}
}

func rollMany(n, pins int) {
	for i := 0; i < n; i++ {
		game.roll(pins)
	}
}

func TestGutterGame(t *testing.T) {
	setUp()
	rollMany(20, 0)
	if game.score() != 0 {
		t.Errorf("Game.score() for a 0 pins game expect 0, got %d", game.score()) 
	}
}

func TestAllOnes(t *testing.T) {
	setUp()
	rollMany(20, 1)
	if game.score() != 20 {
		t.Errorf("Game.score() for all 1 pins rolls expect 20, got %d", game.score())
	}
}
