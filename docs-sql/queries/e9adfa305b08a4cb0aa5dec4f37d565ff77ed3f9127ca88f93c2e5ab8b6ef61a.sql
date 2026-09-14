SELECT year, month, day, count(*) FROM t GROUP BY ROLLUP(year, month, day);
