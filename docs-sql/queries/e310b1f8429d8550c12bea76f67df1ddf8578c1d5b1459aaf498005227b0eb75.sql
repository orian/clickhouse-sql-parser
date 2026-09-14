SELECT finalizeAggregation(uniqThetaUnion(a, b)) AS a_union_b, finalizeAggregation(a) AS a_cardinality, finalizeAggregation(b) AS b_cardinality
FROM
(SELECT arrayReduce('uniqThetaState', [1, 2]) AS a, arrayReduce('uniqThetaState', [2, 3, 4]) AS b);
