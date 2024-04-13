% Caso base: La lista está vacía, no hay elementos para aplanar
aplanar([], []).

% Caso base: Si el primer elemento es una lista, aplanarla y luego aplanar el resto de la lista
aplanar([X|Xs], Zs) :-
    is_list(X),
    aplanar(X, Y),
    aplanar(Xs, Ys),
    append(Y, Ys, Zs).

% Caso recursivo: Si el primer elemento no es una lista, mantenerlo y aplanar el resto de la lista
aplanar([X|Xs], [X|Ys]) :-
    \+ is_list(X),
    aplanar(Xs, Ys).
