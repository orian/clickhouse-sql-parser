WITH histogram(10)(hits) AS hist
SELECT
    round(arrayJoin(hist).1) AS lowerBound,
    round(arrayJoin(hist).2) AS upperBound,
    arrayJoin(hist).3 AS count,
    bar(count, 0, max(count) OVER (), 20) AS b
FROM
(
    SELECT
        path,
        sum(hits) AS hits
    FROM wikistat
    WHERE date(time) = '2015-06-15'
    GROUP BY path
    HAVING hits > 10000
);
