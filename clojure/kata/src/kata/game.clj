(ns kata.game)

(defn get-score [pins]
  (reduce + pins)
)
