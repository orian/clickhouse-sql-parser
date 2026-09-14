SELECT count(*)
FROM hackernews
WHERE hasAnyTokens(lower(comment), 'clickhouse');
