
        CREATE TABLE test_lc (x UInt64, lc LowCardinality(Nullable(String))) ENGINE = MergeTree order by x
    