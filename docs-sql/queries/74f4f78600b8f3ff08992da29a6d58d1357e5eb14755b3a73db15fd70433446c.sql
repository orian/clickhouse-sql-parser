CREATE TABLE example_theta
(
    id UInt32,
    category String
)
ENGINE = Memory;

INSERT INTO example_theta VALUES
(1, 'A'), (2, 'B'), (3, 'A'), (4, 'C'), (5, 'B'), (6, 'A');

SELECT uniqTheta(category) as theta_unique_categories
FROM example_theta;
