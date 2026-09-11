SELECT
    name,
    status,
    element_count,
    formatReadableSize(bytes_allocated) AS size,
    query_count,
    hit_rate,
    found_rate,
    last_exception
FROM system.dictionaries
