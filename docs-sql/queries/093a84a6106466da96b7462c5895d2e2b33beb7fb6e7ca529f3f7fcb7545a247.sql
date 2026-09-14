SELECT
    sum(value) AS S,
    count() AS C,
    S / C
FROM test
