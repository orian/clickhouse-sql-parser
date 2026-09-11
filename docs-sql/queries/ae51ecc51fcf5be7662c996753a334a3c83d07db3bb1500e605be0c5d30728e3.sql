SELECT arrayFill(x, y, z -> x > y AND x < z, [5, 3, 6, 2], [4, 7, 1, 3], [10, 2, 8, 5]) AS res
