SELECT count()
FROM stackoverflow.posts
WHERE (Title != '') AND (Title NOT ILIKE '%clickhouse%') AND (Body NOT ILIKE '%clickhouse%') AND (Id IN (
        SELECT PostId
        FROM stackoverflow.comments
        WHERE Text ILIKE '%clickhouse%'
))
