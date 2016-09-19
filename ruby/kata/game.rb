class Game

  def initialize
    @current_roll = 0
    @rolls = []
  end

  def roll(pins)
    @rolls[@current_roll] = pins
    @current_roll += 1
  end

  def score
    s = 0
    @rolls.each { |pins| s += pins }
    return s
  end

end
