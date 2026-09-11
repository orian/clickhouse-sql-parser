SELECT theilsU(a, b)
FROM (
    SELECT
        number % 10 AS a,
        number % 4 AS b
    FROM
        numbers(150)
);
