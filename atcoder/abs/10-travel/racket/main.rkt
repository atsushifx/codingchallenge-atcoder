#lang racket


;; Initialize
;  const
(define N (expt 10 5))
(define tp0 (hash 't 0 'x 0 'y 0))

; util functions

; solver functions
; read N points from input
(define (readN n)
  (let (
    (v (make-vector (+ n 1) tp0)))
    (begin
      (for ([i n])
        (vector-set! v (+ i 1) (hash 't (read) 'x (read) 'y (read))))

      v
    )))

;;　check point to point can move
(define (check-move tp)
  (let (
    (n (- (vector-length tp) 1))  ;; kukan (not check last point)
    (can-move #t))
    (for ([i n])
      (if (not can-move)
        #f ;; continue
        (begin
          (let (
            (p0 (vector-ref tp i))
            (p1 (vector-ref tp (+ i 1)))
            (dt 0) (dp 0) (d 0)) ;; work 変数
            (begin
              (set! dt (- (hash-ref p1 't) (hash-ref p0 't)))
              (set! dp (+ (abs (- (hash-ref p1 'x) (hash-ref p0 'x)))
                        (abs (- (hash-ref p1 'y) (hash-ref p0 'y)))))
              (set! d (- dt dp))
              (set! can-move (and (>= d 0) (= (remainder d 2) 0))))
          ))))
    can-move  ))


;; main
;
(define (solve)
  (let (
    (n (read))
    (tp 0)) ;; later read
    (begin
      (set! tp (readN n))
      (if (check-move tp) "Yes" "No"))
  ))


; input
;;(define n (read))
;;(define tp (readN n))
;;(if (check-move tp) "Yes"  "No")

(display (solve))
(newline)








