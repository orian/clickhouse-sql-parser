SELECT a, sum(b) FROM t PREWHERE c > 1 WHERE d < 2 GROUP BY a HAVING sum(b) > 10;
