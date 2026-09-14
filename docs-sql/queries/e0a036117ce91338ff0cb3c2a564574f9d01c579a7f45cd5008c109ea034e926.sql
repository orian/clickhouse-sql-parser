SELECT
    formatReadableSize(sum(data_compressed_bytes)) AS total_pinned_compressed,
    formatReadableSize(sum(data_uncompressed_bytes)) AS total_pinned_uncompressed,
    count() AS parts_count
FROM system.snapshot_parts
