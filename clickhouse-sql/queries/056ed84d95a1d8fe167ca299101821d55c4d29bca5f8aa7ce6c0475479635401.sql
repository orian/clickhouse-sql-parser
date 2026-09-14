CREATE TABLE perf_nested_lc_functions
        (
            id UInt64,
            arr_lc Array(LowCardinality(String)),
            map_key_lc Map(LowCardinality(String), String),
            map_value_lc Map(String, LowCardinality(String))
        )
        ENGINE = MergeTree
        ORDER BY id
    