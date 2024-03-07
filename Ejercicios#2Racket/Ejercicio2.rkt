(define (contiene-subcadena? subcadena cadena)
  (regexp-match? (regexp (format "~a" subcadena)) cadena))

(define (filtrar-cadenas subcadena lista-cadenas)
  (filter (lambda (cadena) (contiene-subcadena? subcadena cadena)) lista-cadenas))

(displayln (filtrar-cadenas "la" '("la casa," "el perro," "pintando la cerca,")))

(displayln (filtrar-cadenas "no" '("el no puede ver," "el perro" "no vayas por alla")))

(displayln (filtrar-cadenas "por" '("el arbol esta por alla," "el cepillo," "hablame por mensaje," "por correo," "la tarde,")))

