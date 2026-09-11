SELECT year, month, day, count(*) FROM t GROUP BY year, month, day WITH ROLLUP;
