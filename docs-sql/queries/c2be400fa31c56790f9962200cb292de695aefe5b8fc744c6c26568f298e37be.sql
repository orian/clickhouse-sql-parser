--Si una función de agregación recibe múltiples argumentos, esta función debe aplicarse a múltiples arrays del mismo tamaño.

SELECT arrayReduce('maxIf', [3, 5], [1, 0]);
