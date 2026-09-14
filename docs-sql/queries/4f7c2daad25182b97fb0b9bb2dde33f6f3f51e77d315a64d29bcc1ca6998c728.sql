SELECT
    title,
    like_count,
    dislike_count
FROM youtube
WHERE uploader = 'ClickHouse'
ORDER BY dislike_count DESC;
