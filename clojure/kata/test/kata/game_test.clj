(ns kata.game-test
  (:use clojure.test
    kata.game))

(deftest test-gutter-game
    (is (= 0 (roll (repeat 20 0))))
)
