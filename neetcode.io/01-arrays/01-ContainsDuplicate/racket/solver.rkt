#lang racket
;;
(require racket/contract)
(provide contains-duplicate)

;; solve function
(define/contract (contains-duplicate nums)
  (-> (listof exact-integer?) boolean?)
    (check-dup nums (make-hash))
  )

(define (check-dup nums nhash)
  (cond
    ((null? nums) #f)
    ((hash-has-key? nhash (car nums)) #t)
    (else (begin
      (hash-set! nhash (car nums) #t)
      (check-dup (cdr nums) nhash)
    ))
  ))
