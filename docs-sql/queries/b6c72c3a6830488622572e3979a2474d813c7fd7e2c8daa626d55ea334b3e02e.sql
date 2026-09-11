SELECT * FROM table1 WHERE col1 GLOBAL IN (SELECT col1 FROM table2 WHERE <some_predicate>)
