(define (calcular-monto capital interes anos)
  (cond
    [(= anos 0) capital]
    [else (calcular-monto (+ capital (* capital interes)) interes (- anos 1))]))

(displayln "Capital  Interés  Años  Resultado")
(displayln "---------------------------------")
(displayln (format "2000     0.10     0      ~a" (calcular-monto 2000 0.10 0)))
(displayln (format "2000     0.10     1      ~a" (calcular-monto 2000 0.10 1)))
(displayln (format "2000     0.10     2      ~a" (calcular-monto 2000 0.10 2)))
(displayln (format "2000     0.10     3      ~a" (calcular-monto 2000 0.10 3)))


