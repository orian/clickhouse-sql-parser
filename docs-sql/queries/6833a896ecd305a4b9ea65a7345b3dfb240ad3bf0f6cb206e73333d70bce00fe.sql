CREATE TABLE s_table (x Int8) ENGINE = Log;
INSERT INTO s_table SELECT number FROM numbers(0, 20);
INSERT INTO s_table VALUES (NULL);
SELECT sumCount(x) FROM s_table;
