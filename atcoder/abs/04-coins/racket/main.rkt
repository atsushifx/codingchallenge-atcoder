#lang racket


;; solver
; get k (50yen)
(define (pay-count i j k sum)
  (let ((x (+ (* 500 i) (* 100 j) (* 50 k))))
    (cond
      ((< x sum) #t)
      ((= x sum)
        (begin
          ;(display (list i j k))
          ;(display "\n")
          (set! count (+ count 1))
          #t))
      (else #f)
    )))


(define (loopi i j k sum)
  (if (< i 0)
    #f
    (begin
      (loopi (- i 1) j k sum)
      (loopj i j k sum)
    )))


(define (loopj i j k sum)
  (if (< j 0)
    #f
    (begin
      (loopj i (- j 1) k sum)
      (loopk i j k sum)
      )))

(define (loopk i j k sum)
  (if (< k 0)
    #f
    (begin
      (loopk i j (- k 1) sum)
      (pay-count i j k sum)
      count
    )))


;; input
(define A (read)) ; 500yen
(define B (read)) ; 100yes
(define C (read)) ; 50yes
(define X (read)) ; payment
(define count 0)  ; pay counter

;; test
(loopi A B C X)
