SELECT
    sum(C1) AS C1,
    count(C1) AS C2
FROM
(
    SELECT sum(value) AS C1
    FROM test
    GROUP BY group_index
) AS TBL
