CREATE TABLE example_table (
    id UInt32,
    category String,
    value Float64
) ENGINE = Memory;

INSERT INTO example_table VALUES
(1, 'A', 10.5),
(2, 'B', 20.3),
(3, 'A', 15.7),
(4, 'C', 8.9),
(5, 'B', 12.1),
(6, 'A', 18.4);

SELECT uniq(category) as unique_categories
FROM example_table;
