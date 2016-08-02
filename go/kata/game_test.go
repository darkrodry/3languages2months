package kata

import "testing"

func Test(t *testing.T) {
	var game GameImpl = GameImpl {}
	for i := 0; i < 20; i++ {
		game.roll(0)
	}
}
