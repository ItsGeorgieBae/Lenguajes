% Hechos que representan los platos disponibles en el restaurante junto con su valor calórico y categoría.
plato(guacamole, 200, entrada).
plato(ensalada, 150, entrada).
plato(consome, 300, entrada).
plato(tostadas_caprese, 250, entrada).
plato(pure_de_papa, 350, entrada).

plato(filete_cerdo, 400, carne).
plato(pollo_al_horno, 280, carne).
plato(carne_en_salsa, 320, carne).

plato(tilapia, 160, pescado).
plato(salmon, 300, pescado).
plato(trucha, 225, pescado).

plato(flan, 200, postre).
plato(nueces_con_miel, 500, postre).
plato(naranja_confitada, 450, postre).
plato(flan_de_coco, 375, postre).
plato(fruta_con_helado, 555, postre).

% Hechos que representan los platos principales, que son platos de carne o pescado.
plato_principal(Plato) :-
    plato(Plato, _, carne).
plato_principal(Plato) :-
    plato(Plato, _, pescado).

% Regla que define una comida completa, compuesta por una entrada, un plato principal y un postre.
comida_completa(Entrada, PlatoPrincipal, Postre) :-
    plato(Entrada, _, entrada),
    plato_principal(PlatoPrincipal),
    plato(Postre, _, postre).

% Regla para calcular las calorías totales de una comida completa.
calorias_comida_completa(Entrada, PlatoPrincipal, Postre, CaloriasTotales) :-
    plato(Entrada, CaloriasEntrada, _),
    plato(PlatoPrincipal, CaloriasPlatoPrincipal, _),
    plato(Postre, CaloriasPostre, _),
    CaloriasTotales is CaloriasEntrada + CaloriasPlatoPrincipal + CaloriasPostre.

% Regla para generar las primeras 5 combinaciones de comidas que no superen un maximo de calorías y no repitan ningún elemento en su composición del plato con respecto a las obtenidas con anterioridad.
combinaciones(X, Combinaciones) :-
    findall((Entrada, PlatoPrincipal, Postre), (
        plato(Entrada, _, entrada),
        plato_principal(PlatoPrincipal),
        plato(Postre, _, postre),
        calorias_comida_completa(Entrada, PlatoPrincipal, Postre, Calorias),
        Calorias =< X
    ), CombinacionesSinRepetir),
    generar_combinaciones_sin_repetir(CombinacionesSinRepetir, [], Combinaciones).

% Regla para generar combinaciones sin repetir ningún plato en diferentes combinaciones.
generar_combinaciones_sin_repetir([], _, []).
generar_combinaciones_sin_repetir([(E, PP, P)|T], PlatosSeleccionados, [(E, PP, P)|T2]) :-
    \+ member(E, PlatosSeleccionados),
    \+ member(PP, PlatosSeleccionados),
    \+ member(P, PlatosSeleccionados),
    generar_combinaciones_sin_repetir(T, [E, PP, P|PlatosSeleccionados], T2).
generar_combinaciones_sin_repetir([_|T], PlatosSeleccionados, T2) :-
    generar_combinaciones_sin_repetir(T, PlatosSeleccionados, T2).
    
