
        CREATE TABLE test_lc_query (x UInt64, lc LowCardinality(Nullable(String))) ENGINE = MergeTree order by x
    