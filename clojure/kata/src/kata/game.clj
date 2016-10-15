(ns kata.game)

(defn spare? [roll]
	(= 10 (reduce + (take 2 roll)))
)

(defn get-score [pins]
  (loop [cnt pins acc 0]
  	(if (> (count cnt) 0)
	  	(if (spare? cnt)
	  		(recur (drop 2 cnt) (+ acc (reduce + (take 3 cnt))))
	  		(recur (drop 2 cnt) (+ acc (reduce + (take 2 cnt))))
	  	)
	  	acc
  	)
  )
)