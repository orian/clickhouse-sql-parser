
        SELECT stack, sum(value)
        FROM group_by_low_cardinality_array_join
        ARRAY JOIN stack_array AS stack
        GROUP BY stack
        FORMAT Null
        SETTINGS
            max_threads = 1,
            max_block_size = 65536
    