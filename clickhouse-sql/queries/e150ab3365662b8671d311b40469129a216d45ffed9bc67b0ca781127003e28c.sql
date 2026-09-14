
        SELECT sum(u64), min(payload)
        FROM (SELECT u64, payload FROM sparse_column_index_runs_{run_length})
        ARRAY JOIN range({run_length}) AS element
        FORMAT Null
        SETTINGS enable_lazy_columns_replication = 1
    