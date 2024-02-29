(define (eliminar_elemento elemento lista)
  (define (eliminar-elem e)
    (if (= e elemento) #f e))
  
  (define nueva-lista (map eliminar-elem lista))
  
  (if (member elemento lista)
      (filter identity nueva-lista)
      lista))

(displayln (eliminar_elemento 3 '(1 2 3 4 5)))
(displayln (eliminar_elemento 0 '(1 2 3 4 5)))