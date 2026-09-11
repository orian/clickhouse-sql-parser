CREATE TABLE example_data
(
    id UInt32,
    category String
)
ENGINE = Memory;

INSERT INTO example_data VALUES
(1, 'A'), (2, 'B'), (3, 'A'), (4, 'C'), (5, 'B'), (6, 'A');

SELECT uniqExact(category) as exact_unique_categories
FROM example_data;
