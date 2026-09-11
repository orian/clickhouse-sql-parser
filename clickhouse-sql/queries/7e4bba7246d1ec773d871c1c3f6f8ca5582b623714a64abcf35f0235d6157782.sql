
        SELECT stack_flat, sum(value)
        FROM group_by_low_cardinality_array_join
        GROUP BY stack_flat
        FORMAT Null
        SETTINGS
            max_threads = 1,
            max_block_size = 65536
    