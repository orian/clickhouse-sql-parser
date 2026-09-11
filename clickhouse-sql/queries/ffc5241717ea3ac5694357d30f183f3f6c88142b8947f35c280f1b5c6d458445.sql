
        CREATE TABLE sparse_column_index_runs_{run_length}
        (
            id UInt64,
            u64 UInt64,
            payload String
        )
        ENGINE = MergeTree ORDER BY id
        SETTINGS ratio_of_defaults_for_sparse_serialization = 0.8
    