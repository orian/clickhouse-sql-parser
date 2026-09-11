CREATE TABLE example_hll
(
    id UInt32,
    category String
)
ENGINE = Memory;

INSERT INTO example_hll VALUES
(1, 'A'), (2, 'B'), (3, 'A'), (4, 'C'), (5, 'B'), (6, 'A');

SELECT uniqHLL12(category) AS hll_unique_categories
FROM example_hll;
