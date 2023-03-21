#lang racket
;; libs

;; utility functions
; read number of Nth
(define (readN n)
  (if (<= n 0) '()
    (cons (read) (readN (- n 1)))
  ))

;; solver function
(define (kagamimochi mochi)
  (hash-keys (mochi-hash mochi (make-hash))))


(define (mochi-hash mochi hash)
  (if (null? mochi)
    hash
    (begin
      (hash-set! hash (car mochi) #t)
      (mochi-hash (cdr mochi) hash)
    )
  ))


;; main routin
;  input
(define N (read))
(define mochi (readN N))

; solve
(length (kagamimochi mochi))


