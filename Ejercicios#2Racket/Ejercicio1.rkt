;lista de productos
(define
  ListaProductos
  '(("Arroz" 8 1800)
    ("Frijoles" 5 1200)
    ("Azucar" 6 1100)
    ("cafe" 2 2800)
    ("leche" 9 1200)))

(define (agregarProducto lista nombre cantidad precio)
  (cond ((null? lista)
         (list(list nombre cantidad precio)))
        ((equal? nombre (caar lista))
         (cons (list (caar lista)
                     (+(cadar lista) cantidad)
                     precio)
               (cdr lista)))
               
               
         (else
          (cons (car lista ) (agregarProducto (cdr lista) nombre cantidad precio)))))

(define (venderProducto lista nombre cantidad)
  (cond ((null? lista)
         (display "No existe ese producto vender")
         '())
        ((equal? nombre (caar  lista))
         (cons (list
                (caar lista)
                (- (list-ref (car lista) 1) cantidad)
                (list-ref (car lista) 2))
               (cdr lista)))
        (else
         (cons (car lista)
               (venderProducto (cdr lista) nombre cantidad)
               ))))


;Ejercicio 1
(define (calcularImpuesto factura montoMinimo)
  (define (aplicaImpuesto item)
    (if (>= (* (list-ref item 1) (list-ref item 2)) montoMinimo)
        (* (/ 13 100) (* (list-ref item 1) (list-ref item 2)))
        0))

  (foldl + 0 (map aplicaImpuesto factura)))


;Ejercicio 2
(define (calcularMontoTotal factura)
  (foldl + 0 (map (lambda (item) (* (list-ref item 1) (list-ref item 2))) factura)))



(displayln "Factura #1 ")
(define facturaVenta1
  '(("leche" 3 1200)
    ("Azucar" 1 1100)))

(displayln "Factura de venta:")
(displayln facturaVenta1)
(display "Impuesto total a pagar: ₡")
(display (calcularImpuesto facturaVenta1 2000))
(displayln " ")
(displayln "Monto total de la factura sin impuesto:")
(display (calcularMontoTotal facturaVenta1))
(displayln " ")
(displayln " ")


(displayln "Factura #2 ")
(define facturaVenta2
  '(("Frijoles" 4 1200)
    ("Azucar" 1 1100)))

(displayln "Factura de venta:")
(displayln facturaVenta2)
(display "Impuesto total a pagar: ₡")
(display (calcularImpuesto facturaVenta2 1500))
(displayln " ")
(displayln "Monto total de la factura sin impuesto:")
(display (calcularMontoTotal facturaVenta2))
(displayln " ")
(displayln " ")


(displayln "Factura #3 ")
(define facturaVenta3
  '(("Arroz" 2 1800)
    ("Frijoles" 4 1200)
    ("Azucar" 3 1100)
    ("Cafe" 1 2800)
    ("leche" 2 1200)))

(displayln "Factura de venta:")
(displayln facturaVenta3)
(display "Impuesto total a pagar: ₡")
(display (calcularImpuesto facturaVenta3 2500))
(displayln " ")
(displayln "Monto total de la factura sin impuesto:")
(display (calcularMontoTotal facturaVenta3))

