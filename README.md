# Guía 8

## Ejercicios

### División y Conquista

1. Escribir una función recursiva para determinar si un elemento está en un
    arreglo utilizando búsqueda binaria. Calcule su eficiencia.

2. Programar la busqueda ternaria recursiva. Donde en lugar de dividir el
    arreglo en dos partes iguales, se divide en tres partes iguales. Calcular el
    orden.

3. Se tiene un arreglo de `len(n) >= 3` elementos en forma de pico, esto es:
    estrictamente creciente hasta una determinada posición `p`, y estrictamente
    decreciente a partir de ella (con `0 < p < n-1`). Por ejemplo, en el arreglo
    `[1, 2, 3, 1, 0, -2]` la posición del pico es `p=2`.

    Se pide implementar un algoritmo de división y conquista de orden _O(log n)_
    que encuentre la posición p del pico.

4. Programar una función que realice la multiplicación de dos números enteros utilizando la 
    técnica de división y conquista. Calcular el orden.

5. Sea `a` un vector ordenado de enteros todos distintos de 
    tamaño `n`. Implementar un algoritmo de división y conquista, de complejidad _O(log n)_ en el peor caso, capaz de encontrar un índice `i` tal que `0≤i≤n-1` y `a[i] = i`, suponiendo que tal índice exista.