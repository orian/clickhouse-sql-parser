SELECT
    cramersV(a, b)
FROM
    (
        SELECT
            number % 10 AS a,
            if (number % 12 = 0, (number + 1) % 5, number % 5) AS b
        FROM
            numbers(150)
    );
