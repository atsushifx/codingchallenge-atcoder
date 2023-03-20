#lang racket

;; functions
;; solver
(define (solve)
  (let (
    (a (read)) (b (read))
  )
  (display (if (isEven a b) "Even" "Odd"))
  (display "\n")
  ))

(define (isEven x y)
  (= (* (remainder x 2) (remainder y 2)) 0)
)

;; main routin
(solve)
