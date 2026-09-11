WITH
    (SELECT sum(number) FROM numbers(10)) AS total,
    total * 2 AS doubled,
SELECT total, doubled;
