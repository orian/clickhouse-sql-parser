CREATE TABLE example_table (start Date, end Date) ENGINE = Memory;
INSERT INTO example_table VALUES ('2025-03-03', '2025-03-11'), ('2025-03-06', '2025-03-08'), ('2025-03-07', '2025-03-09'), ('2025-03-11', '2025-03-12');

SELECT start, runningConcurrency(start, end) FROM example_table;
