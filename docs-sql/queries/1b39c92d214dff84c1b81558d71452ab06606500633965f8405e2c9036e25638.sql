SELECT
    cramersV(a, b),
    cramersVBiasCorrected(a, b)
FROM
    (
        SELECT
            number % 10 AS a,
            number % 4 AS b
        FROM
            numbers(150)
    );
