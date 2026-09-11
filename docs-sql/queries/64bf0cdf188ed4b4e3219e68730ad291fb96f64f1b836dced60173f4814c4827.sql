SELECT ll.Date FROM (SELECT * FROM t AS ll LEFT JOIN t1 ON ll.k = t1.k LEFT JOIN t2 ON ll.k = t2.k);
