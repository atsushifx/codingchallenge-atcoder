#lang racket

;; Copyright (c) 2023 Furukawa,Atsushi <atsushifx@aglabo.com>
;;
;; This software is released under the MIT License.
;; https://opensource.org/licenses/MIT

;; solve function
(define (solve)
  (let ([a (read)] [b (read)] [c (read)] [s (read)]) (printf "~s ~s\n" (+ a b c) s)))

;; call solver
(solve)
