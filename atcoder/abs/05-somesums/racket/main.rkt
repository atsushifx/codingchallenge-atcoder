#lang racket

;; utility functions
(define (list->sum L) (apply + L))

;; solver functions
(define (num-list num)
  (let ((hd (quotient num 10)) (tl (modulo num 10)))
    (if (<= hd 0)
      (cons tl null)
      (cons tl (num-list hd))
    )))

;
(define (sum-num num a b)
  (let ((sum (list->sum (num-list num))))
    (if
      (and (>= sum a) (<= sum b))
      sum
      -1 )))

;
(define (make-numbers-list num a b)
  (let ((sum (sum-num num a b)))
    (cond
      ((< num 1) null)
      ((> sum 0) (cons num (make-numbers-list (- num 1) a b)))
      (else (make-numbers-list (- num 1) a b))
      )))

;; solver
(define (solve)
  (let ((nums (make-numbers-list N A B)))

    (list->sum nums)

))


;; input parameter
(define N (read))
(define A (read))
(define B (read))

;; solve
(solve)

