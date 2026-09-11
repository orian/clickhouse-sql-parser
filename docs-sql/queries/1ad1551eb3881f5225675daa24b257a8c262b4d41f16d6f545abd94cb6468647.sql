SELECT arrayFold(
(acc, x, y) -> acc + (x * y),
[1, 2, 3, 4],
[10, 20, 30, 40],
0::Int64
) AS res;
