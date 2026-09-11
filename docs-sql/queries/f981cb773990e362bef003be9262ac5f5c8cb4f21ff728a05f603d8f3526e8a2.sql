SELECT
    JSONExtractString(data, 'type') AS type,
    data
FROM events
