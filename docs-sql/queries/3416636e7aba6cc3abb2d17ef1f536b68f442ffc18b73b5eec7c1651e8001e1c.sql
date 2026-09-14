RESTORE TABLE db.t0, TABLE db.t1
FROM S3('<endpoint>', '<access_key>', '<secret_key>')
SETTINGS
    allow_non_empty_tables = true;
