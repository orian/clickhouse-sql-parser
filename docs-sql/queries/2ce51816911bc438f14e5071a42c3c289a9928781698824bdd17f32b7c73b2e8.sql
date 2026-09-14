RESTORE TABLE db.t0 PARTITIONS '2026-06-01', '2026-06-02', '2026-06-03'
FROM S3('<endpoint>', '<access_key>', '<secret_key>')
SETTINGS
    allow_non_empty_tables = true;
