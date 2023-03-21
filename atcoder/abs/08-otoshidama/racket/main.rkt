#lang racket

;; utility function
(define (nlist->string l)
  (if (null? l)
    ""
    (string-trim (string-append (number->string (car l)) " " (nlist->string (cdr l))))))

;; solver functions
; 10000yen
(define (loop10000yen i j)
  (if (< i 0 )
    #f
    (begin
      (let ((l (loop5000yen i (- N i))))
      (if (list? l)
        l
        (loop10000yen (- i 1) j)
      )))))


; 5000yen
(define (loop5000yen i j)
  (if (< j 0)
    #f
    (begin
      (let
        ( (k (- N i j))
          (sum (+ (* 10000 i) (* 5000 j) (* 1000 (- N i j)))))
        (begin
          ; (printf "~s = ~s yen\n" (list i j k) sum)
          (if (and (>= k 0) (= sum Y))
            (list i j k )
            (loop5000yen i (- j 1))))))))


; solve
(define (solve)
  (let ((l (loop10000yen N N)))
    (if (list? l)
      (begin
        (display (nlist->string l))
        (display "\n")
      )
      (display "-1 -1 -1\n")
    )))

;; main
; input
(define N (read))
(define Y (read))

;; test
(solve)

