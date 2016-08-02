package kata

import "testing"

func TestGutterGame(t *testing.T) {
	var game Game = Game {0}
	for i := 0; i < 20; i++ {
		game.roll(0)
	}
	if game.score() != 0 {
		t.Errorf("Game.score() for a 0 pins game expect 0, got %d", game.score()) 
	}
}

func TestAllOnes(t *testing.T) {
        var game Game = Game {0}
        for i := 0; i < 20; i++ {
                game.roll(1)
        }
        if game.score() != 20 {
                t.Errorf("Game.score() for all 1 pins rolls expect 20, got %d", game.score())
        }
}
