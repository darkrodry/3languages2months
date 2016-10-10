(ns kata.game-test
  (:use clojure.test
    kata.game))

(deftest test-gutter-game
    (is (= 0 (get-score (repeat 20 0))))
)

(deftest test-all-ones
    (is (= 20 (get-score (repeat 20 1))))
)
