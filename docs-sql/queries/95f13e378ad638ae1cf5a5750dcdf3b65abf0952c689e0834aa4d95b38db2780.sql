SELECT key,
       argMax(col1, version) AS col1,
       argMax(col2, version) AS col2
FROM t
WHERE <your predicates>
GROUP BY key;
