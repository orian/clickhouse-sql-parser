SELECT minMappedArrays(a, b)
FROM VALUES('a Array(Int32), b Array(Int64)', ([1, 2], [2, 2]), ([2, 3], [1, 1]));
