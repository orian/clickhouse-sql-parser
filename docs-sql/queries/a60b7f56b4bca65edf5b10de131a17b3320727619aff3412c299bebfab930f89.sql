SELECT key,
       argMax(col1, version) AS col1,
       argMax(col2, version) AS col2
FROM t
WHERE <您的过滤条件>
GROUP BY key;
