
        CREATE TABLE qcc_topk_test (id UInt32, v UInt32, payload FixedString(256)) ENGINE = MergeTree ORDER BY id
        SETTINGS index_granularity = 8192,
                 min_bytes_for_wide_part = 0,
                 min_bytes_for_full_part_storage = 0,
                 add_minmax_index_for_numeric_columns = 0,
                 auto_statistics_types = ''
    