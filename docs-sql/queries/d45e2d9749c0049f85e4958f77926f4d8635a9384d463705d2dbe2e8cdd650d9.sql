CREATE TABLE example_table (a UInt32, b Float32) ENGINE = Memory;
INSERT INTO example_table VALUES (1, 1.001), (2, 1.002), (3, 1.003), (4, 1.004);

SELECT quantileDD(0.01, 0.75)(a), quantileDD(0.01, 0.75)(b) FROM example_table;
