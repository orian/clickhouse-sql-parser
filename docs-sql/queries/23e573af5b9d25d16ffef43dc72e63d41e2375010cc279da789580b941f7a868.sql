SELECT groupArray(*) AS new_array
FROM
(
    SELECT 1
    UNION ALL
    SELECT 2
    UNION ALL
    SELECT 3
)
