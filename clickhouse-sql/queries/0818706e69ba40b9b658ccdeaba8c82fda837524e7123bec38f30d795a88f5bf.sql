
        CREATE TABLE t_streaming_read_round
        (
            value UInt64,
            event_time DateTime64(3),
            PROJECTION commit_order INDEX * TYPE commit_order
        )
        ENGINE = MergeTree ORDER BY tuple()
        SETTINGS enable_block_number_column = 1, enable_block_offset_column = 1, allow_commit_order_projection = 1, merge_selector_algorithm = 'Manual'
    