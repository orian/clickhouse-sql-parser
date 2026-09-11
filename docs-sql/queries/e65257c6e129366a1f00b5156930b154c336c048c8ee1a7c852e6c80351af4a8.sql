--ClickHouse
SELECT arrayJoin(arrayFilter(t -> (t != ''), splitByChar('|', Tags))) AS tags,
        sum(ViewCount) AS views
FROM posts
GROUP BY tags
ORDER BY views DESC
LIMIT 5
