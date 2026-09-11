INSERT INTO dst SELECT
    number + 1 AS key,
    IF(key = 0, 'A', 'B') AS value
FROM numbers(2);

SELECT
    *,
    _part
FROM dst
ORDER BY all;
