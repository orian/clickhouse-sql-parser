SELECT
    `table`,
    name,
    formatReadableSize(sum(data_compressed_bytes)) AS compressed_size,
    formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed_size,
    round(sum(data_uncompressed_bytes) / sum(data_compressed_bytes), 2) AS ratio
FROM system.columns
WHERE (name IN ('Id', 'ViewCount', 'AnswerCount')) AND (`table` IN ('posts_v3', 'posts_v4'))
GROUP BY
    `table`,
    name
ORDER BY
    name ASC,
    `table` ASC
