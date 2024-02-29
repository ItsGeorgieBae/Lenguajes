(define (sub-conjunto? xs ys)
  (cond
    [(null? xs) #t]
    [(null? ys) #f]
    [else (and (member (car xs) ys)
              (sub-conjunto? (cdr xs) (cdr ys)))]))

(sub-conjunto? '() '(a b c d e f)) ; #t
(sub-conjunto? '(a b c) '(a b c d e f)) ; #t
(sub-conjunto? '(a b x) '(a b c d e f)) ; #f