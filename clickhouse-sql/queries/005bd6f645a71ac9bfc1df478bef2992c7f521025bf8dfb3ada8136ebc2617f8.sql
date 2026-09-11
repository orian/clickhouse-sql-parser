
        CREATE TABLE short_agg_test (time DateTime, path String, hits UInt64) ENGINE = MergeTree ORDER BY (path, time)
    