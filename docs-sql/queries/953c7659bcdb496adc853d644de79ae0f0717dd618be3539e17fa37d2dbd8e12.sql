-- A NumericIndexedVector is a binary value, so it is shown with `hex`.
SELECT
    hex(numericIndexedVectorBuild(mapFromArrays([1, 2, 3], [10, 20, 30]))) AS res,
    toTypeName(numericIndexedVectorBuild(mapFromArrays([1, 2, 3], [10, 20, 30]))) AS type;
