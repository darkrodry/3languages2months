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

func rollSpare() {
	game.roll(5)
	game.roll(5)
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

func TestOneSpare(t *testing.T) {
	setUp()
	rollSpare()
	game.roll(3)
	rollMany(17, 0)
	if game.score() != 16 {
		t.Errorf("Game.score() for spare expect 16, got %d", game.score())
	}
}

func TestOneStrike(t *testing.T) {
	setUp()
	game.roll(10)
	game.roll(4)
	game.roll(3)
	rollMany(16, 0)
	if game.score() != 24 {
		t.Errorf("Game.score() for strike expect 24, got %d", game.score())
	}
}
