(ns kata.game)

(defn spare? [roll]
	(= 10 (reduce + (take 2 roll)))
)

(defn strike? [roll]
	(= 10 (first roll))
)

(defn get-score [pins]
  (loop [cnt pins acc 0]
  	(if (> (count cnt) 0)
  		(if (strike? cnt)
  			(recur (drop 1 cnt) (+ acc (reduce + (take 3 cnt))))
		  	(if (spare? cnt)
		  		(recur (drop 2 cnt) (+ acc (reduce + (take 3 cnt))))
		  		(recur (drop 2 cnt) (+ acc (reduce + (take 2 cnt))))
		  	)
	  	)
	  	acc
  	)
  )
)