SELECT
    id,
    title,
    type,
    by,
    time,
    url,
    score
FROM hackernews
WHERE type = 'story'
LIMIT 3
FORMAT Vertical
