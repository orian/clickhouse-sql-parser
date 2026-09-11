SELECT table,
    formatReadableSize(sum(bytes)) as size
    FROM system.parts
    WHERE active
GROUP BY table
