    SELECT
        view_count,
        like_count,
        dislike_count,
        concat('https://youtu.be/', id) AS url,
        title
    FROM youtube
    WHERE (title ILIKE '%ClickHouse%') OR (description ILIKE '%ClickHouse%')
    ORDER BY
        like_count DESC,
        view_count DESC;
