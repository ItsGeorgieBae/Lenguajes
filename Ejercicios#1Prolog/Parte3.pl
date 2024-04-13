% Predicado principal que calcula la distancia de Hamming entre dos cadenas.
distanciaH(X, Y, D) :-
    atom_chars(X, XChars),
    atom_chars(Y, YChars),
    distanciaH_aux(XChars, YChars, 0, D).

% Caso base: Si ambas listas están vacías, la distancia de Hamming es 0.
distanciaH_aux([], [], D, D).

% Si una lista está vacía y la otra no, la distancia de Hamming es la longitud de la lista no vacía.
distanciaH_aux([], [_|Ys], Acc, D) :-
    distanciaH_aux([], Ys, Acc, D).

distanciaH_aux([_|Xs], [], Acc, D) :-
    distanciaH_aux(Xs, [], Acc, D).

% Si ambas listas tienen al menos un elemento.
distanciaH_aux([X|Xs], [Y|Ys], Acc, D) :-
    X \= Y, % Si los elementos son diferentes, incrementamos el contador de distancia.
    NewAcc is Acc + 1,
    distanciaH_aux(Xs, Ys, NewAcc, D).

distanciaH_aux([_|Xs], [_|Ys], Acc, D) :-
    % Si los elementos son iguales, no incrementamos el contador de distancia.
    distanciaH_aux(Xs, Ys, Acc, D).
