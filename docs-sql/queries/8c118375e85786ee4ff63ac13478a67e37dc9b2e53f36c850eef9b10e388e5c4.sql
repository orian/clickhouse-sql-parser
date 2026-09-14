-- The arrayJoin function affects all sections of the query, including the WHERE section. Notice the result 2, even though the subquery returned 1 row.

SELECT sum(1) AS impressions
FROM
(
    SELECT ['Istanbul', 'Berlin', 'Bobruisk'] AS cities
)
WHERE arrayJoin(cities) IN ['Istanbul', 'Berlin'];
