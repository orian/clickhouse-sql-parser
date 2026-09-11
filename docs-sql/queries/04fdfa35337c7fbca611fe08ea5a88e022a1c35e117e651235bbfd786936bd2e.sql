SELECT
    table,
    formatReadableSize(sum(data_compressed_bytes)) AS compressed,
    formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed,
    sum(rows) AS rows
FROM system.parts
WHERE active = 1
  AND database = 'nyc_taxi'
  AND table IN ('trips_small_inferred', 'trips_small_no_pk')
GROUP BY database, table
ORDER BY sum(data_compressed_bytes) DESC;
