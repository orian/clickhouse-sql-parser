SELECT a, b, val FROM t1 INNER JOIN t2 ON t1.a = t2.key OR t1.b = t2.key AND t2.val > 3;
