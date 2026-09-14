SELECT t1.id, t2.name
FROM table1 AS t1 FINAL
INNER JOIN table2 AS t2 FINAL ON t1.id = t2.id;
