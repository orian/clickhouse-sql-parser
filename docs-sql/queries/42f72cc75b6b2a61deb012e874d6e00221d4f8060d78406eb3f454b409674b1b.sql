SET allow_delta_lake_writes = 1;

INSERT INTO my_delta_table
SETTINGS
    delta_lake_insert_max_rows_in_data_file = 1000000,
    delta_lake_insert_max_bytes_in_data_file = 134217728
SELECT * FROM source_table
