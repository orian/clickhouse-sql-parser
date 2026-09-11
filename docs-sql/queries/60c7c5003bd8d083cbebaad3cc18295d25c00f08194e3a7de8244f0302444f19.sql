SELECT finalizeAggregation(uniqThetaIntersect(a, b)) AS a_intersect_b, finalizeAggregation(a) AS a_cardinality, finalizeAggregation(b) AS b_cardinality
FROM
(SELECT arrayReduce('uniqThetaState', [1, 2]) AS a, arrayReduce('uniqThetaState', [2, 3, 4]) AS b);
