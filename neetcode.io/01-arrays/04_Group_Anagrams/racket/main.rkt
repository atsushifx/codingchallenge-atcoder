#lang racket
;; Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
;;
;; This software is released under the MIT License.
;; https://opensource.org/licenses/MIT


(define/contract (group-anagrams strs)
  (-> (listof string?) (listof (listof string?)))
  (let (
    (shash (str-hash strs (make-hash))))
    (hash-values shash)
  ))


(define (str-hash lst shash)
  (if (null? lst)
    shash
    (let (
      (s (car lst))
      (k (sorted-string (car lst)))
      (vs (hash-ref! shash (sorted-string (car lst)) null)))
      (cond
        ((null? vs) (set! vs (list s)))
        (else (set! vs (cons s vs))))
      (hash-set! shash k vs)
      (str-hash (cdr lst) shash)
      shash
    )))

(define (sorted-string str)
  (list->string (sort (string->list str) char<?))
)

;; main roution
(group-anagrams '("eat" "tea" "tan" "ate" "nat" "bat"))


