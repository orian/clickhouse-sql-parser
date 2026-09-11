--Se uma função de agregação recebe múltiplos argumentos, ela deve ser aplicada a múltiplos arrays de mesmo tamanho.

SELECT arrayReduce('maxIf', [3, 5], [1, 0]);
