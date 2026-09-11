    SELECT
       by,
       count() AS comments
    FROM hackernews
    WHERE (type IN ('story', 'comment')) AND ((title ILIKE '%ClickHouse%') OR (text ILIKE '%ClickHouse%'))
    GROUP BY by
    ORDER BY comments DESC
    LIMIT 5
