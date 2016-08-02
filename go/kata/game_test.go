package kata

import "testing"

func Test(t *testing.T) {
	var game Game = Game {}
	for i := 0; i < 20; i++ {
		game.roll(0)
	}
}
