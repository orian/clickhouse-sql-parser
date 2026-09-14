SELECT finalizeAggregation(uniqThetaNot(a, b)) AS a_not_b, finalizeAggregation(a) AS a_cardinality, finalizeAggregation(b) AS b_cardinality
FROM
(SELECT arrayReduce('uniqThetaState', [2, 3, 4]) AS a, arrayReduce('uniqThetaState', [1, 2]) AS b);
