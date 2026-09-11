SELECT maxMappedArrays(a, b)
FROM VALUES('a Array(Char), b Array(Int64)', (['x', 'y'], [2, 2]), (['y', 'z'], [3, 1]));
