#lang racket

;; utility functions

; read number of Nth
(define (readN n)
  (if (<= n 0) '()
    (cons (read) (readN (- n 1)))
  ))


;; for solving
; list divide
(define (shifter list)
  (if (null? list) '()
    (let (
      (hd (car list))
      (tl (shifter (cdr list))))
    (cond
      ((> (modulo hd 2) 0) #f)
      ((not tl) #f)
      (else (cons (/ hd 2) tl))
    ))
  ))

; count list shift
(define (shift-count list cnt)
  (let
    ((shiftedList (shifter list))
    )
    (begin
      (if (not shiftedList) cnt
        (shift-count shiftedList (+ cnt 1))
         ))))

;; solve main
(define (solve list)
  (shift-count list 0)
)


;; main
(define N (read))
(define al (readN N))
(solve al)


