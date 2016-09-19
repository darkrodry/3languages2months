require 'test/unit'
require_relative 'game' 

class GameTest < Test::Unit::TestCase

  def test_gutter_game
    game = Game.new
    20.times do
      game.roll(0)
    end
    assert_equal 0, game.score
  end

  def test_all_ones
    game = Game.new
    20.times do
      game.roll(1)
    end
    assert_equal 20, game.score
  end
end
