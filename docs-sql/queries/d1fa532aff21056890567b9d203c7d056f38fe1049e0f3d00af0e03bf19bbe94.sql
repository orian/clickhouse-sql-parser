SELECT
    `table`,
    formatReadableSize(data_compressed_bytes) AS compressed_size,
    formatReadableSize(data_uncompressed_bytes) AS uncompressed_size
FROM system.data_skipping_indices
WHERE startsWith(`table`, 'otel_logs')
