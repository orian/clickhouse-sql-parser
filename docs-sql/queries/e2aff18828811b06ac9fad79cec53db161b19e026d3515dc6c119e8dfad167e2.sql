EXPLAIN indexes = 1
SELECT count(*)
FROM hackernews
WHERE hasAnyTokens(lower(comment), 'clickhouse');
