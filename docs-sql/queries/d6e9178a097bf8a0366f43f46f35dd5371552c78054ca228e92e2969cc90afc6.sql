SELECT *
FROM deltaLake('s3://my-bucket/warehouse/ga4_events/')
SETTINGS
    delta_lake_snapshot_start_version = 42,
    delta_lake_snapshot_end_version = 47
