package kata

import "testing"

func TestGutterGame(t *testing.T) {
	var game Game = Game {}
	for i := 0; i < 20; i++ {
		game.roll(0)
	}
	if game.score() != 0 {
		t.Errorf("Game.score() for a 0 pins game expect 0, got %d", game.score()) 
	}
}
