
        INSERT INTO iceberg_manifest_entries SELECT number FROM numbers(100000)
        SETTINGS allow_experimental_insert_into_iceberg = 1, iceberg_insert_max_rows_in_data_file = 1,
                 max_block_size = 1, min_insert_block_size_rows = 1, min_insert_block_size_bytes = 1
    