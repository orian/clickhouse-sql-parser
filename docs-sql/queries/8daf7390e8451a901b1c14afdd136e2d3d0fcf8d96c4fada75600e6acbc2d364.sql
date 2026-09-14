--If an aggregate function takes multiple arguments, then this function must be applied to multiple arrays of the same size.

SELECT arrayReduce('maxIf', [3, 5], [1, 0]);
