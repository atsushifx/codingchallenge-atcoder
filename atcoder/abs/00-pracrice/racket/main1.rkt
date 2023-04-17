#lang racket

;; input
(define a (read))
(define b (read))
(define c (read))
(define s (read))

;; output
(printf "~s ~s\n" (+ a b c) s)
