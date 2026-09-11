SELECT
arrayIntersect([1, 2], [1, 3], [2, 3]) AS empty_intersection,
arrayIntersect([1, 2], [1, 3], [1, 4]) AS non_empty_intersection
