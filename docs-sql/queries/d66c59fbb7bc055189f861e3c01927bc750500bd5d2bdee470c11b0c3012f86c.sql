SELECT count() FROM my_iceberg_table
SETTINGS iceberg_metadata_log_level = 'manifest_file_entry';

SYSTEM FLUSH LOGS iceberg_metadata_log;

SELECT content_type, file_path, pruning_status
FROM system.iceberg_metadata_log
WHERE query_id = '<previous_query_id>';
