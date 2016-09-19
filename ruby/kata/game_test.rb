require 'test/unit'
require_relative 'game' 

class GameTest < Test::Unit::TestCase
  def test_gutter_game
    game = Game.new
    20.times do
      game.roll(0)
    end
  end
end
