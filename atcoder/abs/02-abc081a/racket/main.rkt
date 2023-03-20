#lang racket

;; utility functions
; sum of list (numbers)
(define (list->sum L) (apply + L))

; bind char list to int list
(define (map->to_i list)
  (map
    (lambda (ch) (- (char->integer ch) (char->integer #\0)))
  list)
)


;; solver
(define (solve)
  (let (
    (marbles
      (map->to_i (string->list (number->string (read))))
      ))
  (list->sum marbles)
  ))

;; main
(solve)






