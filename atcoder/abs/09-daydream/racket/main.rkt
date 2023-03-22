#lang racket

;; main
; input
(define s (symbol->string (read)))

; solve
(define matcher #rx"^(dream|dreamer|erase|eraser)*$")

(display (if (list? (regexp-match matcher s)) "YES\n" "NO\n"))
