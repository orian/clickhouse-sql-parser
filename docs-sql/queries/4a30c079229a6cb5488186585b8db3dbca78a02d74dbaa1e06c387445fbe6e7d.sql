SELECT
    histogram(10)(hits) AS hist
FROM
(
    SELECT
        path,
        sum(hits) AS hits
    FROM wikistat
    WHERE date(time) = '2015-06-15'
    GROUP BY path
    HAVING hits > 10000
)
FORMAT Vertical;
