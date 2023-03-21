#lang racket

;; utility functions

; read number of Nth
(define (readN n)
  (if (<= n 0) '()
    (cons (read) (readN (- n 1)))
  ))


;; solver function
(define (game cards)
  (if (null? cards)
    0
    (let (
      (turn (length cards))
      (point (car cards))
      (sum (game (cdr cards)))
    )
    (begin
      ;(display (list turn point sum))
      ;(display "\n")
      (if (> (remainder turn 2) 0)
        (+ sum point)
        (- sum point)))
    )))


;; input param
(define N (read))
(define cards (readN N))

;; main
(game (sort cards <))

