
        CREATE TABLE test_lc_insert (k1 LowCardinality(String), k2 LowCardinality(String)) ENGINE = MergeTree order by (k1, k2)
    