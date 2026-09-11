
        SELECT *
        FROM
        (
            SELECT
                RegionID,
                toDateTime(EventDate) AS date,
                count() AS quantity
            FROM test.hits
            GROUP BY
                RegionID,
                date
        ) AS regions
        ASOF LEFT JOIN
        (
            SELECT
                RegionID,
                EventTime
            FROM test.hits
        ) AS ids ON (regions.RegionID = ids.RegionID) AND (regions.date < ids.EventTime)
    