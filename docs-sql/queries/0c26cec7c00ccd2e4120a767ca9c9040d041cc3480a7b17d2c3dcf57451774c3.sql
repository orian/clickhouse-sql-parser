SELECT year, month, day, count(*) FROM t GROUP BY CUBE(year, month, day);
