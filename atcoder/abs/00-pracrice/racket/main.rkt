#lang racket

;; solver
(define (solve)
  (let (
      (a (read)) (b (read)) (c (read))
      (s (read)))
    (begin
      (display (+ a b c))
      (display " ")
      (display s)
      (display "\n")
      )
  ))

;; main routin
(solve)

