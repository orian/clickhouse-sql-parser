
        INSERT INTO group_by_low_cardinality_array_join
        WITH
            if(number < 8190, number, intHash64(number) % 8190) AS stack_number,
            bitOr(
                bitShiftLeft(toUInt128(cityHash64(stack_number)), 64),
                toUInt128(stack_number)) AS stack
        SELECT
            [stack] AS stack_array,
            stack AS stack_flat,
            toUInt64(1) AS value
        FROM numbers(25000000)
        SETTINGS
            max_threads = 1,
            max_insert_threads = 1,
            max_block_size = 25000000,
            max_insert_block_size = 25000000,
            min_insert_block_size_rows = 25000000,
            min_insert_block_size_bytes = 0
    