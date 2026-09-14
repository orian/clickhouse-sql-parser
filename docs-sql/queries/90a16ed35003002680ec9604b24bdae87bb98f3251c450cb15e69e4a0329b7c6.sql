CREATE TABLE fl_interval (id String, start Float32, end Float32) ENGINE = Memory;
INSERT INTO fl_interval VALUES ('a', 1.1, 2.9), ('a', 2.5, 3.2), ('a', 4, 5);

SELECT id, intervalLengthSum(start, end), toTypeName(intervalLengthSum(start, end)) FROM fl_interval GROUP BY id ORDER BY id;
