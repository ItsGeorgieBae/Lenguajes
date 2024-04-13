% Caso base: Si la primera lista está vacía, siempre es un subconjunto
sub_conjunto([], _).
% Caso recursivo: Verificar si el primer elemento de la primera lista está en la segunda lista,
% y luego verificar el resto de la lista
sub_conjunto([X|Xs], Ys) :-
    member(X, Ys),
    sub_conjunto(Xs, Ys).
