WITH (value) -> value + 1 AS increment
SELECT increment(first_result) AS second_increment
FROM
(
    SELECT increment(number) AS first_result
    FROM numbers(3)
);
