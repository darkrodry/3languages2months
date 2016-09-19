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
    score = 0
    for frame in (0..19).step(2) do  
      if is_spare?(frame)
        score += 10 + @rolls[frame+2]
      else
        score += @rolls[frame] + @rolls[frame+1]
      end
    end
    return score
  end

  def is_spare?(frame)
    return @rolls[frame] + @rolls[frame+1] == 10
  end
end
